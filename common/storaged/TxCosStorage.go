package storaged

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go-zero-dandan/common/resd"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// 检查是否实现了工厂接口
var _ InterfaceFactory = (*TxCosProvider)(nil)
var _ InterfaceStorage = (*TxCosStorage)(nil)

// TxCosProvider 腾讯云文件管理
type TxCosProvider struct {
	config *ProviderConfig
	client *cos.Client
}

// TxCosStorage 实现文件管理器接口
type TxCosStorage struct {
	config *ProviderConfig
	client *cos.Client
	baseUploader
}

// Init 初始化操作
func (t *TxCosProvider) Init() error {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(t.config.Endpoint)
	b := &cos.BaseURL{BucketURL: u}
	txCosClient := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: os.Getenv(t.config.Secret), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: os.Getenv(t.config.Key), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})
	t.client = txCosClient
	return nil
}

// CreateDownloader 创建文件上传器
func (t *TxCosProvider) CreateDownloader(downloaderConfig *DownloaderConfig) (InterfaceStorage, error) {
	return &TxCosStorage{
		config: t.config,
		client: t.client,
	}, nil
}

// CreateUploader 创建文件下载器
func (t *TxCosProvider) CreateUploader(uploaderConfig *UploaderConfig) (InterfaceStorage, error) {
	if uploaderConfig == nil {
		return nil, resd.NewErr("uploaderConfig未配置")
	}
	if uploaderConfig.FileType == "" {
		return nil, resd.NewErr("uploaderConfig的FileType未提供文件类型")
	}
	uploader := &TxCosStorage{
		config: t.config,
		client: t.client,
	}
	if uploaderConfig == nil || uploaderConfig.MaxMemorySize == 0 {
		uploader.MaxMemorySize = defaultConfig[uploaderConfig.FileType].MaxMemorySize
	}
	if uploaderConfig == nil || uploaderConfig.MaxFileSize == 0 {
		uploader.MaxFileSize = defaultConfig[uploaderConfig.FileType].MaxFileSize
	}
	if uploaderConfig == nil || len(uploaderConfig.FileMimeAccept) == 0 {
		uploader.AcceptMimes = defaultConfig[uploaderConfig.FileType].AcceptMimes
	}
	uploader.Result = &UploadResult{}
	return uploader, nil
}

// GetHash 获取文件sha1哈希值
func (t *TxCosStorage) GetHash(r *http.Request, formKey string) (string, error) {
	return t.getHash(r, formKey)
}

// Upload 简单上传文件
func (t *TxCosStorage) Upload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {

	return nil, nil
}

// MultipartUpload 分片上传文件
func (t *TxCosStorage) MultipartUpload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {

	return nil, nil
}

// MultipartDownload 分片下载文件
func (t *TxCosStorage) MultipartDownload(w http.ResponseWriter, path string) (err error) {

	return nil
}

// UploadImg 上传图片，提供图片专属处理参数
func (t *TxCosStorage) UploadImg(r *http.Request, config *UploadImgConfig) (res *UploadResult, err error) {
	t.Type = FileTypeImage //图片上传方法，强制存储类型为图片
	t.Request = r          //传递请求参数，以免下载方法中需要使用
	// 根据form key获取文件
	if err = t.processFileGet(); err != nil {
		return nil, err
	}
	// 获取文件大小和校验
	if err = t.processFileSize(); err != nil {
		return nil, err
	}
	// 获取文件格式和校验
	if err = t.processFileType(); err != nil {
		return nil, err
	}
	// 获取文件哈希值
	if err = t.processFileHash(); err != nil {
		return nil, err
	}
	//拼接存储目录路径，个人习惯，图片放在img文件夹下
	objectName := fmt.Sprintf("img/%s/%s%s", getDirName(), t.Result.Hash, t.Result.Ext)
	if err = t.upload(objectName); err != nil {
		return nil, err
	}
	return t.Result, nil
}

// Download 下载文件
func (t *TxCosStorage) Download(w http.ResponseWriter, objectName string, saveFileName ...string) error {
	//调用腾讯云接口获取文件内容
	resObject, err := t.client.Object.Get(context.Background(), objectName, nil)
	if err != nil {
		return resd.Error(err)
	}
	defer resObject.Body.Close()
	_, err = io.Copy(w, resObject.Body)
	if err != nil {
		return resd.Error(err)
	}
	saveName := ""
	if len(saveFileName) > 0 {
		saveName = saveFileName[0]
	} else {
		saveName = filepath.Base(objectName)
	}
	// 设置响应头，允许web端获取Content-Disposition头信息查看文件名
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", url.PathEscape(saveName)))
	w.Header().Set("Content-Type", "text/plain")
	return nil
}

// upload 上传的具体实现
func (t *TxCosStorage) upload(objectName string) (err error) {
	_, err = t.client.Object.Put(context.Background(), objectName, t.File, nil)
	if err != nil {
		return resd.Error(err)
	}
	t.Result.Path = objectName
	t.Result.Url = t.config.Endpoint + "/" + objectName
	return nil
}
