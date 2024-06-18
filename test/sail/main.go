package main

import (
	"fmt"
	"github.com/HYY-yu/sail-client"
	"time"
)

type Config struct {
	Name string
	Host string
	Port string
	Mode string

	Database string

	UserRpc struct {
		Etcd struct {
			Hosts []string
			Key   string
		}
	}
	Redisx struct {
		Host string
		Pass string
	}
	JwtAuth struct {
		AccessSecret string
	}
}

func main() {
	var cfg Config

	s := sail.New(&sail.MetaConfig{
		ETCDEndpoints:  "127.0.0.1:2379",
		ProjectKey:     "2-public",
		Namespace:      "plat",
		Configs:        "plat-api.yaml",
		ConfigFilePath: "./conf",
		LogLevel:       "DEBUG",
	}, sail.WithOnConfigChange(func(configFileKey string, s *sail.Sail) {
		if s.Err() != nil {
			fmt.Println(s.Err())
			return
		}

		fmt.Println(s.Pull())

		v, err := s.MergeVipers()
		if err != nil {
			fmt.Println(err)
			return
		}
		v.Unmarshal(&cfg)
		fmt.Println(cfg, "\n", cfg.Database)
	}))
	if s.Err() != nil {
		fmt.Println(s.Err())
		return
	}

	fmt.Println(s.Pull())

	v, err := s.MergeVipers()
	if err != nil {
		fmt.Println(err)
		return
	}
	v.Unmarshal(&cfg)
	fmt.Println(cfg, "\n", cfg.Database)

	for {
		time.Sleep(time.Second)
	}
}
