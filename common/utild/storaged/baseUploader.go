package storaged

import (
	"errors"
	"go-zero-dandan/common/utild"
	"mime/multipart"
	"net/http"
)

type baseUploader struct {
	MaxSize        int64
	Request        *http.Request
	File           multipart.File
	FileHeader     *multipart.FileHeader
	Type           FileType
	Hash           string
	FileMimeAccept []string
}
type defaultConfigStruct struct {
	MaxSize     int64
	AcceptMimes []string
}

var defaultConfig = map[FileType]defaultConfigStruct{
	FileTypeImage: {MaxSize: 5 << 20, AcceptMimes: []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif",
	}},
	FileTypeFile: {MaxSize: 100 << 20, AcceptMimes: []string{
		"image/jpeg", "image/jpg", "image/png", "image/gif",
	}},
}

func (t *baseUploader) processFileGet(r *http.Request) (err error) {
	switch t.Type {
	case FileTypeImage:
		_ = t.Request.ParseMultipartForm(20 << 20) //20M,控制表单数据在内存中的存储大小，超过该值，则会自动将表单数据写入磁盘临时文件
		t.File, t.FileHeader, err = r.FormFile("img")
	case FileTypeFile:
		_ = t.Request.ParseMultipartForm(200 << 20) //200M
		t.File, t.FileHeader, err = r.FormFile("file")
	default:
		err = errors.New("unsupported file in file get")
	}

	return err
}
func (t *baseUploader) processFileSize() (err error) {
	if t.FileHeader.Size > t.MaxSize {
		return errors.New("file size limited")
	}
	return nil
}
func (t *baseUploader) processFileHash() (err error) {
	//获取文件hash
	t.Hash, err = utild.GetFileHashHex(t.File)
	if err != nil {
		return err
	}
	//重新指向文件头，避免上传minio时长度不对
	_, err = t.File.Seek(0, 0)
	return err
}
func (t *baseUploader) processFileType() (err error) {
	var validImageTypes map[string]bool
	for _, v := range t.FileMimeAccept {
		validImageTypes[v] = true
	}

	// 读取文件前 512 字节
	buffer := make([]byte, 512)
	if _, err = t.File.Read(buffer); err != nil {
		return errors.New("unsupport image type")
	}
	// 判断文件 MIME 类型是否为图片类型
	mime := http.DetectContentType(buffer)
	if !validImageTypes[mime] {
		return errors.New("invalid img type")
	}
	//重新指向文件头，避免后续操作问题
	if _, err = t.File.Seek(0, 0); err != nil {
		return err
	}
	return nil
}
