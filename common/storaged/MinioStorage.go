package storaged

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go-zero-dandan/common/resd"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
)

// 检查是否实现了工厂接口
var _ InterfaceFactory = (*MinioProvider)(nil)
var _ InterfaceStorage = (*MinioStorage)(nil)

// MinioProvider 实现文件管理渠道工厂
type MinioProvider struct {
	config *ProviderConfig
	client *minio.Client
}

// MinioStorage 实现文件管理器接口
type MinioStorage struct {
	config *ProviderConfig
	client *minio.Client
	baseUploader
}

// Init 初始化操作
func (t *MinioProvider) Init() error {
	minioClient, err := minio.New(t.config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(t.config.Key, t.config.Secret, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	t.client = minioClient
	return nil
}

// CreateDownloader 创建文件上传器
func (t *MinioProvider) CreateDownloader(downloaderConfig *DownloaderConfig) (InterfaceStorage, error) {
	return &MinioStorage{
		config: t.config,
		client: t.client,
	}, nil
}

// CreateUploader 创建文件下载器
func (t *MinioProvider) CreateUploader(uploaderConfig *UploaderConfig) (InterfaceStorage, error) {
	if uploaderConfig == nil {
		return nil, resd.NewErr("uploaderConfig未配置")
	}
	if uploaderConfig.FileType == "" {
		return nil, resd.NewErr("uploaderConfig的FileType未提供文件类型")
	}
	uploader := &MinioStorage{
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
func (t *MinioStorage) GetHash(r *http.Request, formKey string) (string, error) {
	return t.getHash(r, formKey)
}

// Upload 简单上传文件
func (t *MinioStorage) Upload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {

	return nil, nil
}

// MultipartUpload 分片上传文件
func (t *MinioStorage) MultipartUpload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {

	return nil, nil
}

// MultipartDownload 分片下载文件
func (t *MinioStorage) MultipartDownload(w http.ResponseWriter, path string) (err error) {

	return nil
}

// UploadImg 上传图片，提供图片专属处理参数
func (t *MinioStorage) UploadImg(r *http.Request, config *UploadImgConfig) (res *UploadResult, err error) {
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
	//上传文件
	if err = t.upload(objectName); err != nil {
		return nil, err
	}
	return t.Result, nil
}

// Download 下载文件
func (t *MinioStorage) Download(w http.ResponseWriter, objectName string, saveFileName ...string) error {
	bucketName := t.config.Bucket
	//调用minio接口获取文件内容
	object, err := t.client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return resd.Error(err)
	}
	defer object.Close()
	//虽然object是结构体，但在使用io.Copy时他会自动调用Reader处理
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
func (t *MinioStorage) upload(objectName string) (err error) {
	bucketName := t.config.Bucket
	_, err = t.client.PutObject(context.Background(), bucketName, objectName, t.File, t.FileHeader.Size,
		minio.PutObjectOptions{ContentType: "binary/octet-stream"})
	if err != nil {
		return resd.Error(err)
	}
	t.Result.Path = objectName
	t.Result.Url = "http://" + t.config.Endpoint + "/" + bucketName + "/" + objectName
	return nil
}
