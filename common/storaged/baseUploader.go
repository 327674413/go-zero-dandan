package storaged

import (
	"errors"
	"go-zero-dandan/common/imgd"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

// UploadResult 上传结果集
type UploadResult struct {
	Hash     string
	Name     string
	Mime     string
	Ext      string
	SizeByte int64
	SizeText string
	Url      string
	Path     string
}

// baseUploader 上传基础类
type baseUploader struct {
	MaxFileSize   int64
	MaxMemorySize int64
	Request       *http.Request
	File          multipart.File
	FileHeader    *multipart.FileHeader
	Type          FileType
	AcceptMimes   []string
	LocalPath     string
	Result        *UploadResult
}

// defaultConfigStruct 默认配置
type defaultConfigStruct struct {
	MaxFileSize   int64
	MaxMemorySize int64
	AcceptMimes   []string
}

// 定义默认配置
var defaultConfig = map[FileType]defaultConfigStruct{
	FileTypeImage: {MaxFileSize: 5 << 20, MaxMemorySize: 10 << 20, AcceptMimes: []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif",
	}},
	FileTypeFile: {MaxFileSize: 100 << 20, MaxMemorySize: 200 << 20, AcceptMimes: []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif",
	}},
}

// processFileGet 根据上传器类型，获取对应文件，目前写死图片img，文件file，视频video等
func (t *baseUploader) processFileGet() (err error) {
	switch t.Type {
	case FileTypeImage:
		_ = t.Request.ParseMultipartForm(20 << 20) //20M,控制表单数据在内存中的存储大小，超过该值，则会自动将表单数据写入磁盘临时文件
		t.File, t.FileHeader, err = t.Request.FormFile("img")
	case FileTypeFile:
		_ = t.Request.ParseMultipartForm(200 << 20) //200M
		t.File, t.FileHeader, err = t.Request.FormFile("file")
	default:
		err = errors.New("unsupported file in file get")
	}

	return err
}

// processFileSize 校验及获取文件大小信息
func (t *baseUploader) processFileSize() (err error) {
	if t.FileHeader.Size > t.MaxFileSize {
		return resd.Fail("file size limited", resd.UploadFileSizeLimited1, utild.FormatFileSize(t.MaxFileSize))
	}
	t.Result.SizeByte = t.FileHeader.Size
	t.Result.Name = t.FileHeader.Filename
	t.Result.Ext = strings.ToLower(filepath.Ext(t.FileHeader.Filename))
	t.Result.SizeText = utild.FormatFileSize(t.FileHeader.Size)
	return nil
}

// processFileHash 获取文件哈希
func (t *baseUploader) processFileHash() (err error) {
	//获取文件hash
	t.Result.Hash, err = utild.GetFileHashHex(t.File)
	if err != nil {
		return err
	}
	//重新指向文件头，避免上传minio时长度不对
	_, err = t.File.Seek(0, 0)
	t.Result.Name = t.FileHeader.Filename
	return err
}

// processImg 图片处理
func (t *baseUploader) processImg(config *UploadImgConfig) (err error) {
	if config == nil {
		return nil
	}
	imager, err := imgd.NewImg(t.Result.Path)
	if err != nil {
		return err
	}
	//图片缩放
	if config.Resize != nil {

		switch config.Resize.Type {
		case UploadImgResizeTypeCover:
			imager.ResizeCover(config.Resize.Width, config.Resize.Height)
		case UploadImgResizeTypeContain:
			imager.ResizeContain(config.Resize.Width, config.Resize.Height)
		case UploadImgResizeTypeFill:
			imager.ResizeFill(config.Resize.Width, config.Resize.Height)
		case UploadImgResizeTypeWidthFix:
			imager.ResizeWidthFix(config.Resize.Width)
		case UploadImgResizeTypeHeightFix:
			imager.ResizeHeightFix(config.Resize.Height)
		default:
			return resd.NewErr("不支持的图片缩放方式")
		}
	}
	//图片压缩
	if config.Quality > 0 {
		//暂无压缩方案，imaging的压缩png会变大
	}
	//水印
	if config.Watermark != nil {
		if config.Watermark.Type == imgd.WatermarkTypeImg {
			if config.Watermark.Path == "" {
				return resd.NewErr("图片水印请传入Path")
			}
			imager.WatermarkImg(config.Watermark)

		}
	}
	return imager.Output(t.Result.Path)
}
func (t *baseUploader) getHash(r *http.Request, formKey string) (string, error) {
	file, _, err := r.FormFile(formKey)
	if err != nil {
		return "", resd.Error(err)
	}
	//获取文件hash
	hash, err := utild.GetFileHashHex(file)
	if err != nil {
		return "", resd.Error(err)
	}
	//重新指向文件头，避免上传minio时长度不对
	_, err = file.Seek(0, 0)
	if err != nil {
		return "", resd.Error(err)
	}
	return hash, nil
}
func (t *baseUploader) processFileType() (err error) {
	validImageTypes := make(map[string]bool)
	for _, v := range t.AcceptMimes {
		validImageTypes[v] = true
	}

	// 读取文件前 512 字节
	buffer := make([]byte, 512)
	if _, err = t.File.Read(buffer); err != nil {
		return errors.New("unsupport image type")
	}
	// 判断文件 MIME 类型是否为图片类型
	mime := http.DetectContentType(buffer)
	if _, ok := validImageTypes[mime]; !ok {
		return resd.Fail("invalid img type", resd.UploadImageTypeLimited1, t.GetLimitedExtStr())
	}
	//重新指向文件头，避免后续操作问题
	if _, err = t.File.Seek(0, 0); err != nil {
		return err
	}
	t.Result.Mime = mime
	return nil
}
func (t *baseUploader) GetLimitedExtStr() string {
	return strings.Join(t.AcceptMimes, ",")
}

func (t *baseUploader) ResizeImg() {

}
