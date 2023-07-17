package biz

import (
	"go-zero-dandan/app/asset/api/internal/types"
	"go-zero-dandan/common/utild/storaged"
	"net/http"
)

type Upload struct {
}

func (t *Upload) Upload(r *http.Request, conf *storaged.StorageConfig) (*types.UploadResp, error) {
	storageFactory, err := storaged.NewStorage(conf)
	if err != nil {
		return nil, err
	}
	hash, err := storageFactory.GetHash()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
