package storaged

import (
	"fmt"
	"go-zero-dandan/common/resd"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// 检查是否实现了接口
var _ InterfaceUploader = (*LocalUploader)(nil)
var _ InterfaceStorage = (*LocalStorage)(nil)

// LocalStorage 本地文件管理
type LocalStorage struct {
	config *StorageConfig
	svc    *StorageSvc
}
type LocalUploader struct {
	config *StorageConfig
	baseUploader
}

func (t *LocalStorage) CreateUploader(uploaderConfig *UploaderConfig) (InterfaceUploader, error) {
	if uploaderConfig == nil {
		return nil, resd.NewErr("uploaderConfig未配置")
	}
	if uploaderConfig.FileType == "" {
		return nil, resd.NewErr("uploaderConfig的FileType未提供文件类型")
	}
	uploader := &LocalUploader{}
	uploader.config = t.config
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
func (t *LocalStorage) Init() error {
	if t.config.LocalPath == "" {
		return resd.NewErr("本地存储时LocalPath必传")
	}
	return nil

}
func (t *LocalUploader) GetHash(r *http.Request, formKey string) (string, error) {
	return t.getHash(r, formKey)
}
func (t *LocalUploader) UploadImg(r *http.Request) (res *UploadResult, err error) {
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
	dirName := getDirName()
	dirPath := filepath.Join(t.config.LocalPath, "img", dirName)
	if err = os.MkdirAll(dirPath, 0755); err != nil {
		return nil, err
	}
	if err = t.upload(dirPath); err != nil {
		return nil, err
	}
	return t.Result, nil
}
func (t *LocalStorage) UploadFile() {

}

func (t *LocalUploader) upload(dirPath string) (err error) {
	//拼接返回的url地址
	url := ""
	if t.Request.TLS == nil {
		url = "http://"
	} else {
		url = "https://"
	}
	url = url + t.Request.Host
	//根据雪花id生成新的文件名
	newFileName := fmt.Sprintf("%s%s", t.Result.Hash, t.Result.Ext)
	//获取完整的存储路径
	savePath := path.Join(dirPath, newFileName)
	//存储文件
	tempFile, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()
	io.Copy(tempFile, t.File)
	t.Result.Url = url + "/" + savePath
	return nil
}
