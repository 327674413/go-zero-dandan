package storaged

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go-zero-dandan/common/resd"
	"net/http"
	"net/url"
	"os"
)

// 检查是否实现了工厂接口
var _ InterfaceUploader = (*MinioUploader)(nil)
var _ InterfaceStorage = (*TxCosStorage)(nil)

// TxCosStorage 腾讯云文件管理
type TxCosStorage struct {
	config *StorageConfig
	client *cos.Client
}
type TxCosUploader struct {
	config *StorageConfig
	client *cos.Client
	baseUploader
}

func (t *TxCosStorage) Init() error {
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
func (t *TxCosStorage) CreateUploader(uploaderConfig *UploaderConfig) (InterfaceUploader, error) {
	if uploaderConfig == nil {
		return nil, resd.NewErr("uploaderConfig未配置")
	}
	if uploaderConfig.FileType == "" {
		return nil, resd.NewErr("uploaderConfig的FileType未提供文件类型")
	}
	uploader := &TxCosUploader{
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
func (t *TxCosUploader) UploadImg(r *http.Request, config *UploadImgConfig) (res *UploadResult, err error) {
	t.Type = FileTypeImage
	t.Request = r
	if err = t.processFileGet(); err != nil {
		return nil, err
	}
	if err = t.processFileSize(); err != nil {
		return nil, err
	}
	if err = t.processFileType(); err != nil {
		return nil, err
	}
	if err = t.processFileHash(); err != nil {
		return nil, err
	}
	//判断与生成目录
	objectName := fmt.Sprintf("img/%s/%s%s", getDirName(), t.Result.Hash, t.Result.Ext)
	if err = t.upload(objectName); err != nil {
		return nil, err
	}
	return t.Result, nil
}
func (t *TxCosUploader) Download(r *http.Request, pathAndFileName string) error {

	return nil
}
func (t *TxCosUploader) GetHash(r *http.Request, formKey string) (string, error) {
	return t.getHash(r, formKey)
}
func (t *TxCosUploader) upload(objectName string) (err error) {
	_, err = t.client.Object.Put(context.Background(), objectName, t.File, nil)
	if err != nil {
		return resd.Error(err)
	}
	t.Result.Url = t.config.Endpoint + "/" + objectName
	return nil
}
