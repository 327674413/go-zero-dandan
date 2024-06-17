package configServer

import (
	"encoding/json"
	"github.com/HYY-yu/sail-client"
)

type Config struct {
	ETCDEndpoints  string `toml:"etcd_endpoints"` // 逗号分隔的ETCD地址，0.0.0.0:2379,0.0.0.0:12379,0.0.0.0:22379
	ProjectKey     string `toml:"project_key"`
	Namespace      string `toml:"namespace"`
	Configs        string `toml:"configs"`
	ConfigFilePath string `toml:"config_file_path"` // 本地配置文件存放路径，空代表不存储本都配置文件
	LogLevel       string `toml:"log_level"`        // 日志级别(DEBUG\INFO\WARN\ERROR)，默认 WARN
}

type Sail struct {
	*sail.Sail
	sail.OnConfigChange
	c *Config
}

func NewSail(cfg *Config) *Sail {
	//s := sail.New(&sail.MetaConfig{
	//	ETCDEndpoints:  cfg.ETCDEndpoints,
	//	ProjectKey:     cfg.ProjectKey,
	//	Namespace:      cfg.Namespace,
	//	Configs:        cfg.Configs,
	//	ConfigFilePath: cfg.ConfigFilePath,
	//	LogLevel:       cfg.LogLevel,
	//})
	return &Sail{c: cfg}
}

func (s *Sail) Build() error {
	var opts []sail.Option

	if s.OnConfigChange != nil {
		opts = append(opts, sail.WithOnConfigChange(s.OnConfigChange))
	}

	s.Sail = sail.New(&sail.MetaConfig{
		ETCDEndpoints:  s.c.ETCDEndpoints,
		ProjectKey:     s.c.ProjectKey,
		Namespace:      s.c.Namespace,
		Configs:        s.c.Configs,
		ConfigFilePath: s.c.ConfigFilePath,
		LogLevel:       s.c.LogLevel,
	}, opts...)
	return s.Sail.Err()
}

func (s *Sail) FromJsonBytes() ([]byte, error) {
	if err := s.Pull(); err != nil {
		return nil, err
	}

	return s.fromJsonBytes(s.Sail)
}
func (s *Sail) fromJsonBytes(sail *sail.Sail) ([]byte, error) {
	v, err := sail.MergeVipers()
	if err != nil {
		return nil, err
	}
	data := v.AllSettings()
	return json.Marshal(data)
}
func (s *Sail) Error() error {
	return nil
}

//func (s *Sail) SetOnChange(f OnChange) {
//	s.OnConfigChange = func(configFileKey string, sail *sail.Sail) {
//		data, err := s.fromJsonBytes(sail)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		if err = f(data); err != nil {
//			fmt.Println("OnChange err ", err)
//		}
//
//	}
//}
