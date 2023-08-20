package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

type Command struct {
	Path   string
	Tables []string
}
type Config struct {
	GoctlModelUrl string `yaml:"GoctlModelUrl"`
}

var config Config
var commands = []Command{
	{Path: path.Join(rootPath, "user/model"), Tables: []string{"user_crony", "user_info", "user_main", "user_union"}},
}

const rootPath = "/Users/yelin/go_dev/project/src/go-zero-dandan/app"

func main() {
	// 读取 YAML 文件内容
	yamlFile, err := os.ReadFile("../cmd-dev.yml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v \n", err)
	}
	// 解析 YAML 数据

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to parse YAML data: %v \n", err)
	}

	for _, cmd := range commands {
		for _, table := range cmd.Tables {
			err = updateModelFile(cmd.Path, table)
			if err != nil {
				fmt.Printf("------------- Error: Path: %s , Table:%s , Update Fail:%v -----------", cmd.Path, table, err)
			} else {
				fmt.Printf("------------- Table:%s , Updated ----------- \n", table)
			}
		}

	}
}

const dbAddr = ""

func updateModelFile(path, table string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}
	cmd := fmt.Sprintf("goctl model mysql datasource --ignore-columns=\"delete_at\" -url=\"%s\" -table=\"%s\" . -style goZero -home  ../../../common/goctl/1.5.0", config.GoctlModelUrl, table)
	var command *exec.Cmd
	if runtime.GOOS == "windows" {
		command = exec.Command("cmd", "/C", cmd)
	} else {
		command = exec.Command("/bin/sh", "-c", cmd)
	}

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err = command.Run()
	if err != nil {
		return err
	}
	return nil
}
