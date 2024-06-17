package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-dandan/app/user/model"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	DB string `yaml:"DB"`
}

var config Config

func main() {
	// 读取 YAML 文件内容
	yamlFile, err := os.ReadFile("./conf.yml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v \n", err)
	}
	// 解析 YAML 数据

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to parse YAML data: %v \n", err)
	}
	db := sqlx.NewMysql(config.DB)
	data := make([]*model.UserCrony, 0)
	//var data []model.UserCrony
	//var data model.UserCrony
	err = db.QueryRowsPartialCtx(context.Background(), &data, "select * from staff")
	if err != nil {
		logx.Error(err)
		return
	}
	fmt.Println(data)
}
