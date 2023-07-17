package storaged

import (
	"errors"
	"go-zero-dandan/common/resd"
	"go-zero-dandan/common/utild"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

type UploadResult struct {
	Hash     string
	Name     string
	Mime     string
	Ext      string
	SizeByte int64
	SizeText string
	Url      string
}
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
type defaultConfigStruct struct {
	MaxFileSize   int64
	MaxMemorySize int64
	AcceptMimes   []string
}

var defaultConfig = map[FileType]defaultConfigStruct{
	FileTypeImage: {MaxFileSize: 5 << 20, MaxMemorySize: 10 << 20, AcceptMimes: []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif",
	}},
	FileTypeFile: {MaxFileSize: 100 << 20, MaxMemorySize: 200 << 20, AcceptMimes: []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif",
	}},
}

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
func (t *baseUploader) processFileSize() (err error) {
	if t.FileHeader.Size > t.MaxFileSize {
		return resd.Fail("file size limited", resd.UploadFileSizeLimited1, utild.FormatFileSize(t.MaxFileSize))
	}
	t.Result.SizeByte = t.FileHeader.Size
	t.Result.Name = t.FileHeader.Filename
	t.Result.Ext = filepath.Ext(t.FileHeader.Filename)
	t.Result.SizeText = utild.FormatFileSize(t.FileHeader.Size)
	return nil
}
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
