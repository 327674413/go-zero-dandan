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
)

// 检查是否实现了工厂接口
var _ InterfaceUploader = (*MinioUploader)(nil)
var _ InterfaceStorage = (*MinioStorage)(nil)

// MinioStorage Minio文件管理
type MinioStorage struct {
	config *StorageConfig
	client *minio.Client
}
type MinioUploader struct {
	config *StorageConfig
	client *minio.Client
	baseUploader
}

func (t *MinioStorage) Init() error {
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
func (t *MinioStorage) CreateUploader(uploaderConfig *UploaderConfig) (InterfaceUploader, error) {
	if uploaderConfig == nil {
		return nil, resd.NewErr("uploaderConfig未配置")
	}
	if uploaderConfig.FileType == "" {
		return nil, resd.NewErr("uploaderConfig的FileType未提供文件类型")
	}
	uploader := &MinioUploader{
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
func (t *MinioUploader) UploadImg(r *http.Request, config *UploadImgConfig) (res *UploadResult, err error) {
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
func (t *MinioUploader) Download(w http.ResponseWriter, objectName string) error {
	bucketName := t.config.Bucket
	object, err := t.client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return resd.Error(err)
	}
	defer object.Close()
	/*
		// 获取文件信息
		fileInfo, err := object.Stat()
		if err != nil {
			return err
		}
	*/
	_, err = io.Copy(w, object)
	if err != nil {
		return resd.Error(err)
	}

	// 设置响应头
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	//w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", url.PathEscape("蛋蛋.png")))
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Type", "text/plain")
	return nil
}
func (t *MinioUploader) GetHash(r *http.Request, formKey string) (string, error) {
	return t.getHash(r, formKey)
}
func (t *MinioUploader) upload(objectName string) (err error) {
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
