package storaged

import (
	"net/http"
)

// 检查是否实现了工厂接口
var _ InterfaceStorage = (*LocalStorage)(nil)

// LocalStorage 本地上传
type LocalStorage struct {
	config *StorageConfig
	baseUploader
}

func (t *LocalStorage) UploadImg(r *http.Request) (err error) {
	t.Type = FileTypeImage
	if err = t.processFileGet(r); err != nil {
		return err
	}
	if err = t.processFileSize(); err != nil {
		return err
	}
	if err = t.processFileType(); err != nil {
		return err
	}
	if err = t.processFileHash(); err != nil {
		return err
	}

	return nil
}
func (t *LocalStorage) GetHash() (string, error) {
	err := t.processFileHash()
	if err != nil {
		return "", nil
	}

	return t.Hash, nil
}
func (t *LocalStorage) UploadFile() {

}
