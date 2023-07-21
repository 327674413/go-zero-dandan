package storaged

// NewStorage 文件工厂入口
func NewStorage(config *StorageConfig) (storage InterfaceStorage, err error) {
	if config == nil {
		panic("未传入文件管理config配置")
	}
	switch config.Provider {
	case ProviderLocal:
		storage = &LocalStorage{config: config}
	case ProviderMinio:
		storage = &MinioStorage{config: config}
	case ProviderTxCos:
		storage = &TxCosStorage{config: config}
	case ProviderAliOss:
		storage = &AliOssStorage{config: config}
	default:
		panic("暂不支持的文件管理类型")
	}
	if err = storage.Init(); err != nil {
		return nil, err
	}
	return storage, nil

}
