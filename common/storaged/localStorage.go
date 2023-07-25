package storaged

import (
	"fmt"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

// 检查是否实现了接口
var _ InterfaceFactory = (*LocalProvider)(nil)
var _ InterfaceStorage = (*LocalStorage)(nil)

// LocalProvider 实现文件管理渠道工厂
type LocalProvider struct {
	config *ProviderConfig
	svc    *StorageSvc
}

// LocalStorage 实现文件管理器接口
type LocalStorage struct {
	config *ProviderConfig
	baseUploader
}

// Init 初始化操作
func (t *LocalProvider) Init() error {
	if t.config.LocalPath == "" {
		return resd.NewErr("本地存储时LocalPath必传")
	}
	return nil

}

// CreateDownloader 创建文件上传器
func (t *LocalProvider) CreateDownloader(downloaderConfig *DownloaderConfig) (InterfaceStorage, error) {
	return &LocalStorage{
		config: t.config,
	}, nil
}

// CreateUploader 创建文件下载器
func (t *LocalProvider) CreateUploader(uploaderConfig *UploaderConfig) (InterfaceStorage, error) {
	if uploaderConfig == nil {
		return nil, resd.NewErr("uploaderConfig未配置")
	}
	if uploaderConfig.FileType == "" {
		return nil, resd.NewErr("uploaderConfig的FileType未提供文件类型")
	}
	uploader := &LocalStorage{
		config: t.config,
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
func (t *LocalStorage) GetHash(r *http.Request, formKey string) (string, error) {
	return t.getHash(r, formKey)
}

// Upload 简单上传文件
func (t *LocalStorage) Upload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {

	return nil, nil
}

// MultipartUpload 分片上传文件
func (t *LocalStorage) MultipartUpload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {

	return nil, nil
}

// MultipartDownload 分片下载文件
func (t *LocalStorage) MultipartDownload(w http.ResponseWriter, path string) (err error) {

	return nil
}

// UploadImg 上传图片，提供图片专属处理参数
func (t *LocalStorage) UploadImg(r *http.Request, config *UploadImgConfig) (res *UploadResult, err error) {
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
	dirName := getDirName()                                      //获取年-月-日格式的目录ing成
	dirPath := filepath.Join(t.config.LocalPath, "img", dirName) //拼接存储目录路径，个人习惯，图片放在img文件夹下
	//确保目录存在
	if err = os.MkdirAll(dirPath, 0755); err != nil {
		return nil, err
	}
	//上传文件
	if err = t.upload(dirPath); err != nil {
		return nil, err
	}
	//对文件进行图片相关的处理
	if err = t.processImg(config); err != nil {
		return nil, err
	}
	return t.Result, nil
}

// Download 下载文件
func (t *LocalStorage) Download(w http.ResponseWriter, objectName string, saveFileName ...string) error {
	file, err := os.Open(objectName) //根据路径，打开文件
	if err != nil {
		return resd.Error(err)
	}
	defer file.Close()
	_, err = io.Copy(w, file) //通过copy方式，写入请求体中，该方式为流式下载
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
func (t *LocalStorage) upload(dirPath string) (err error) {
	//拼接返回的url地址
	url := utild.GetRequestDomain(t.Request)
	//根据雪花id生成新的文件名
	newFileName := fmt.Sprintf("%s%s", t.Result.Hash, t.Result.Ext)
	//获取完整的存储路径
	savePath := path.Join(dirPath, newFileName)
	if err != nil {
		return resd.Error(err)
	}
	//存储文件
	tempFile, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()
	//将文件内容写入存储
	io.Copy(tempFile, t.File)
	t.Result.Path = savePath
	t.Result.Url = url + "/" + savePath
	return nil
}
