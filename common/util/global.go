package util

import (
	"gopkg.in/yaml.v3"
	"os"
)

type globalObject struct {
	DbPrefix string `yaml:"DbPrefix"` //数据库表前缀

}

// GlobalObj 定义一个全局的对外对象
var GlobalObj *globalObject

// Reload 从配置文件去加载参数
func (t *globalObject) Reload() {
	yamlFile, err := os.ReadFile("config/zinx.yaml")
	if err != nil {
		panic(err)
	}

	// 解析 YAML 文件
	err = yaml.Unmarshal(yamlFile, &GlobalObj)
	if err != nil {
		panic(err)
	}

}

// init 提供一个init方法,被导包时会在main方法之前自动执行，且多次导包只会执行一次
func init() {
	//如果没配置时，使用默认配置值
	//GlobalObj = &globalObject{}
	//GlobalObj.Reload()
}
