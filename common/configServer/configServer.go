package configServer

import (
	"errors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrNotSetConfig = errors.New("未设置配置信息")

type ConfigServer interface {
	FromJsonBytes() ([]byte, error)
	Error() error
}
type configServer struct {
	ConfigServer
	configFile string
}

func NewConfigServer(configFile string, s ConfigServer) *configServer {
	return &configServer{
		ConfigServer: s,
		configFile:   configFile,
	}
}
func (s *configServer) MustLoad(v any) error {
	if s.ConfigServer.Error() == nil {
		logx.Info("1")
		return s.ConfigServer.Error()
	}
	if s.configFile == "" && s.ConfigServer == nil {
		logx.Info("2")
		return ErrNotSetConfig
	}
	if s.ConfigServer == nil {
		logx.Info("3")
		// 使用gozero默认方式
		conf.MustLoad(s.configFile, v)
	}
	data, err := s.ConfigServer.FromJsonBytes()
	if err != nil {
		return err
	}
	return conf.LoadFromJsonBytes(data, v)
}
