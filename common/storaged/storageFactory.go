package storaged

import (
	"errors"
	"go-zero-dandan/common/resd"
)

// NewStorage 文件工厂入口
func NewStorage(config *StorageConfig) (storage InterfaceStorage, err error) {
	if config == nil {
		return nil, errors.New("storage config required")
	}
	switch config.Provider {
	case ProviderLocal:
		storage = &LocalStorage{config: config}
	case ProviderMinio:
		storage = &MinioStorage{config: config}
	default:
		return nil, resd.NewErr("暂不支持的文件管理类型")
	}
	if err = storage.Init(); err != nil {
		return nil, err
	}
	return storage, nil

}
