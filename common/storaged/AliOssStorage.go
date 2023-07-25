package storaged

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-zero-dandan/common/resd"
	"io"
	"net/http"
	"net/url"
)

// 检查是否实现了工厂接口
var _ InterfaceUploader = (*AliOssUploader)(nil)
var _ InterfaceStorage = (*AliOssStorage)(nil)

// AliOssStorage 阿里云文件管理
type AliOssStorage struct {
	config *StorageConfig
	client *oss.Client
}
type AliOssUploader struct {
	config *StorageConfig
	client *oss.Client
	baseUploader
}

func (t *AliOssStorage) Init() error {
	client, err := oss.New("https://"+t.config.Endpoint, t.config.Key, t.config.Secret)
	if err != nil {
		panic(err)
	}
	t.client = client
	return nil
}
func (t *AliOssStorage) CreateUploader(uploaderConfig *UploaderConfig) (InterfaceUploader, error) {
	if uploaderConfig == nil {
		return nil, resd.NewErr("uploaderConfig未配置")
	}
	if uploaderConfig.FileType == "" {
		return nil, resd.NewErr("uploaderConfig的FileType未提供文件类型")
	}
	uploader := &AliOssUploader{
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
func (t *AliOssUploader) UploadImg(r *http.Request, config *UploadImgConfig) (res *UploadResult, err error) {
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
func (t *AliOssUploader) Download(w http.ResponseWriter, objectName string) error {
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
	// 设置响应头
	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	//w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", url.PathEscape("蛋蛋.png")))
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Type", "text/plain")
	return nil
}
func (t *AliOssUploader) GetHash(r *http.Request, formKey string) (string, error) {
	return t.getHash(r, formKey)
}
func (t *AliOssUploader) upload(objectName string) (err error) {
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
