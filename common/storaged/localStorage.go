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
	"strconv"
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
func (t *LocalStorage) Upload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {
	t.Type = FileTypeFile //文件上传方法，强制存储类型文件
	t.Request = r         //传递请求参数，以免下载方法中需要使用(目前好像没啥用了)
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
	//普通上传直接按年月日目录
	dirName := GetDateDir()
	dirPath := filepath.Join(t.config.LocalPath, t.Bucket, t.DirName, dirName)
	//确保目录存在，744通常用于普通文件
	if err = os.MkdirAll(dirPath, 0744); err != nil {
		return nil, err
	}
	//上传文件
	if err = t.upload(dirPath); err != nil {
		return nil, err
	}

	return t.Result, nil
}

// MultipartUpload 分片上传文件
func (t *LocalStorage) MultipartUpload(r *http.Request, config *UploadConfig) (res *UploadResult, err error) {
	if config.FileSha1 == "" {
		return nil, resd.NewErr("分片上传必须提供文件sha1", resd.MultipartUploadFileHashRequired)
	}
	t.Type = FileTypeMultipart //文件上传方法，强制存储类型文件
	t.Request = r              //传递请求参数，以免下载方法中需要使用
	// 根据form key获取文件
	if err = t.processFileGet(); err != nil {
		return nil, err
	}
	// 按照文件哈希 + 分片索引作为文件名
	fname := fmt.Sprintf("%s_%d", config.FileSha1, config.ChunkIndex)
	//分片上传放到专属目录，并根据哈希值前2位进行分目录，避免文件太多
	dirPath := filepath.Join(t.config.LocalPath, t.Bucket, t.DirName, config.FileSha1[:2])
	//确保目录存在，744通常用于普通文件
	if err = os.MkdirAll(dirPath, 0744); err != nil {
		return nil, err
	}
	//上传文件
	if err = t.uploadMultipart(dirPath + "/" + fname); err != nil {
		return nil, err
	}
	return t.Result, nil
}

// MultipartMerge 分片上传合并
func (t *LocalStorage) MultipartMerge(fileSha1 string, saveName string, chunkCount int) (*UploadResult, error) {
	// 合并后的文件路径0
	mergedFilePath := filepath.Join(t.config.LocalPath, t.Bucket, t.DirName, GetDateDir(), fmt.Sprintf("%s%s", fileSha1, filepath.Ext(saveName)))
	err := os.MkdirAll(path.Dir(mergedFilePath), 0744)
	if err != nil {
		return nil, resd.Error(err)
	}

	mergedFile, err := os.Create(mergedFilePath)
	if err != nil {
		return nil, resd.Error(err)
	}
	defer mergedFile.Close()
	// 读取每个分块文件数据并加入到合并文件中
	for i := 0; i < chunkCount; i++ {
		chunkFilePath := filepath.Join(t.config.LocalPath, t.Bucket, t.DirName, fileSha1[:2], fileSha1+"_"+strconv.Itoa(i)) // 分块文件路径
		chunkData, err := os.ReadFile(chunkFilePath)
		if err != nil {
			return nil, resd.Error(err, resd.MergeFileChunkNotFound)

		}

		_, err = mergedFile.Write(chunkData)
		if err != nil {
			return nil, resd.Error(err)
		}
	}
	mergeSha1, err := utild.GetFileSha1ByOsFile(mergedFile)
	if err != nil {
		return nil, resd.Error(err)
	}
	if mergeSha1 != fileSha1 {
		return nil, resd.NewErr("合并后文件sha1不相等")
	}
	// 删除已合并的分块文件
	for i := 0; i < chunkCount; i++ {
		chunkFilePath := filepath.Join(t.config.LocalPath, t.Bucket, t.DirName, fileSha1[:2], fileSha1+"_"+strconv.Itoa(i)) // 分块文件路径
		if err != nil {
			return nil, resd.Error(err)

		}
		err = os.Remove(chunkFilePath)
		if err != nil {
			return nil, resd.Error(err)
		}
	}
	// 获取文件大小
	mergedFileInfo, err := os.Stat(mergedFilePath)
	if err != nil {
		return nil, resd.Error(err)
	}

	t.Result.Path = mergedFilePath
	t.Result.Name = saveName
	t.Result.Url = "http://" + t.config.Endpoint + "/" + saveName
	t.Result.SizeByte = mergedFileInfo.Size()
	t.Result.SizeText = utild.FormatFileSize(t.Result.SizeByte)
	t.Result.Mime, err = t.GetFileMimeByOsFile(mergedFile)
	if err != nil {
		return nil, resd.Error(err)
	}
	return t.Result, nil
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
		return nil, resd.Error(err)
	}
	// 获取文件大小和校验
	if err = t.processFileSize(); err != nil {
		return nil, resd.Error(err)
	}
	// 获取文件格式和校验
	if err = t.processFileType(); err != nil {
		return nil, resd.Error(err)
	}
	// 获取文件哈希值
	if err = t.processFileHash(); err != nil {
		return nil, resd.Error(err)
	}
	dirName := GetDateDir()                                                    //获取年-月-日格式的目录ing成
	dirPath := filepath.Join(t.config.LocalPath, t.Bucket, t.DirName, dirName) //拼接存储目录路径，个人习惯，图片放在img文件夹下
	//确保目录存在
	if err = os.MkdirAll(dirPath, 0755); err != nil {
		return nil, resd.Error(err)
	}
	//上传文件
	if err = t.upload(dirPath); err != nil {
		return nil, resd.Error(err)
	}
	//对文件进行图片相关的处理
	if err = t.processImg(config); err != nil {
		return nil, resd.Error(err)
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

// upload 普通上传的具体实现
func (t *LocalStorage) upload(dirPath string) (err error) {
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
	t.Result.Url = "http://" + t.config.Endpoint + "/" + savePath
	return nil
}

// uploadMultipart 分片上传的具体实现
func (t *LocalStorage) uploadMultipart(savePath string) (err error) {
	//存储文件
	tempFile, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer tempFile.Close()
	//将文件内容写入存储
	io.Copy(tempFile, t.File)
	t.Result.Path = savePath
	return nil
}
