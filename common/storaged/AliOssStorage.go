package storaged

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-zero-dandan/common/resd"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
)

// 检查是否实现了工厂接口
var _ InterfaceFactory = (*AliOssProvider)(nil)
var _ InterfaceStorage = (*AliOssStorage)(nil)

// AliOssProvider 阿里云文件管理
type AliOssProvider struct {
	config *ProviderConfig
	client *oss.Client
}

// AliOssStorage 实现文件管理器接口
type AliOssStorage struct {
	config *ProviderConfig
	client *oss.Client
	baseUploader
}

// Init 初始化操作
func (t *AliOssProvider) Init() error {
	client, err := oss.New("https://"+t.config.Endpoint, t.config.Key, t.config.Secret)
	if err != nil {
		panic(err)
	}
	t.client = client
	return nil
}

// CreateDownloader 创建文件上传器
func (t *AliOssProvider) CreateDownloader(downloaderConfig *DownloaderConfig) (InterfaceStorage, error) {
	return &AliOssStorage{
		config: t.config,
		client: t.client,
	}, nil
}

// CreateUploader 创建文件下载器
func (t *AliOssProvider) CreateUploader(uploaderConfig *UploaderConfig) (InterfaceStorage, error) {
	if uploaderConfig == nil {
		return nil, resd.NewErr("uploaderConfig未配置")
	}
	if uploaderConfig.FileType == "" {
		return nil, resd.NewErr("uploaderConfig的FileType未提供文件类型")
	}
	uploader := &AliOssStorage{
		config: t.config,
		client: t.client,
	}
	if uploaderConfig == nil || uploaderConfig.MaxMemorySize == 0 {
		uploader.MaxMemorySize = defaultConfig[uploaderConfig.FileType].MaxMemorySize
	} else {
		uploader.MaxMemorySize = uploaderConfig.MaxMemorySize
	}
	if uploaderConfig == nil || uploaderConfig.MaxFileSize == 0 {
		uploader.MaxFileSize = defaultConfig[uploaderConfig.FileType].MaxFileSize
	} else {
		uploader.MaxFileSize = uploaderConfig.MaxFileSize
	}
	if uploaderConfig == nil || len(uploaderConfig.AcceptMimes) == 0 {
		uploader.AcceptMimes = defaultConfig[uploaderConfig.FileType].AcceptMimes
	} else {
		uploader.AcceptMimes = uploaderConfig.AcceptMimes
	}
	if uploaderConfig == nil || len(uploaderConfig.RejectMimes) == 0 {
		uploader.RejectMimes = defaultConfig[uploaderConfig.FileType].RejectMimes
	} else {
		uploader.RejectMimes = uploaderConfig.RejectMimes
	}
	if uploaderConfig == nil || uploaderConfig.DirName == "" {
		uploader.DirName = defaultConfig[uploaderConfig.FileType].DirName
	} else {
		uploader.DirName = uploaderConfig.DirName
	}
	if uploaderConfig == nil || uploaderConfig.FormKey == "" {
		uploader.FormKey = defaultConfig[uploaderConfig.FileType].FormKey
	} else {
		uploader.FormKey = uploaderConfig.FormKey
	}
	if uploaderConfig == nil || uploaderConfig.Bucket == "" {
		uploader.Bucket = t.config.Bucket
	} else {
		uploader.Bucket = uploaderConfig.Bucket
	}
	uploader.Result = &UploadResult{}
	return uploader, nil
}

// Upload 简单上传文件
func (t *AliOssStorage) Upload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {

	return nil, nil
}

// MultipartUpload 分片上传文件
func (t *AliOssStorage) MultipartUpload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {

	return nil, nil
}

// MultipartMerge 分片上传合并
func (t *AliOssStorage) MultipartMerge(fileSha1 string, saveName string, chunkCount int) (*UploadResult, error) {
	return nil, nil
}

// MultipartDownload 分片下载文件
func (t *AliOssStorage) MultipartDownload(w http.ResponseWriter, path string) (err error) {

	return nil
}

// UploadImg 上传图片，提供图片专属处理参数
func (t *AliOssStorage) UploadImg(r *http.Request, config *UploadImgConfig) (res *UploadResult, err error) {
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
	objectName := fmt.Sprintf("img/%s/%s%s", GetDateDir(), t.Result.Hash, t.Result.Ext)
	if err = t.upload(objectName); err != nil {
		return nil, err
	}
	return t.Result, nil
}

// Download 下载文件
func (t *AliOssStorage) Download(w http.ResponseWriter, objectName string, saveFileName ...string) error {
	//调用阿里云接口获取文件内容
	bucket, err := t.client.Bucket(t.config.Bucket)
	if err != nil {
		return resd.Error(err)
	}
	object, err := bucket.GetObject(objectName, nil)
	if err != nil {
		return resd.Error(err)
	}
	defer object.Close()
	_, err = io.Copy(w, object)
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
func (t *AliOssStorage) upload(objectName string) (err error) {
	bucket, err := t.client.Bucket(t.config.Bucket)
	if err != nil {
		return resd.Error(err)
	}
	//err = bucket.PutObjectFromFile(objectName, "文件路径")
	err = bucket.PutObject(objectName, t.File)
	if err != nil {
		return resd.Error(err)
	}
	t.Result.Path = objectName
	t.Result.Url = "https://" + t.config.Bucket + "." + t.config.Endpoint + "/" + objectName
	return nil
}
