package storaged

import "errors"

// NewStorage 文件工厂入口
func NewStorage(config *StorageConfig) (InterfaceStorage, error) {
	if config == nil {
		return nil, errors.New("storage config required")
	}
	switch config.Provider {
	case ProviderLocal:
		return &LocalStorage{config: config}, nil
	}

	return nil, errors.New("Unsupported Provider")
}
