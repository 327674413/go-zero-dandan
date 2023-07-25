package storaged

// NewProvider 文件工厂入口
func NewProvider(config *ProviderConfig) (provider InterfaceFactory, err error) {
	if config == nil {
		panic("未传入文件管理config配置")
	}
	switch config.Provider {
	case ProviderLocal:
		provider = &LocalProvider{config: config}
	case ProviderMinio:
		provider = &MinioProvider{config: config}
	case ProviderTxCos:
		provider = &TxCosProvider{config: config}
	case ProviderAliOss:
		provider = &AliOssProvider{config: config}
	default:
		panic("暂不支持的文件管理渠道")
	}
	if err = provider.Init(); err != nil {
		return nil, err
	}
	return provider, nil

}
