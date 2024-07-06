package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const projectPath = "/Users/yelin/go_dev/project/src/go-zero-dandan"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./gen.sh {api service_name [-prod] | rpc service_name [-prod] | model [-prod]}")
		return
	}

	command := os.Args[1]
	var serviceName string
	var prodFlag string

	if command != "model" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: ./gen.sh {api service_name [-prod] | rpc service_name [-prod] | model [-prod]}")
			return
		}
		serviceName = os.Args[2]
		if len(os.Args) == 4 {
			prodFlag = os.Args[3]
		}
	}

	switch command {
	case "api":
		runApiCommand(serviceName, prodFlag)
	case "rpc":
		runRpcCommand(serviceName, prodFlag)
	case "model":
		runModelCommand(prodFlag)
	default:
		fmt.Println("Usage: ./gen.sh {api service_name [-prod] | rpc service_name [-prod] | model [-prod]}")
	}
}

func runApiCommand(serviceName, prodFlag string) {
	if prodFlag == "-prod" {
		apiFiles, err := getApiFiles()
		if err != nil {
			fmt.Printf("Error finding .api files: %v\n", err)
			return
		}
		runCommand("goctl", append([]string{"api", "go"}, append(apiFiles, "-dir", "./", "-style", "goZero", "-home", filepath.Join(projectPath, "common", "goctl", "1.5.0"))...)...)
	} else {
		buildGoctl()
		changeDirectory(filepath.Join(projectPath, "app", serviceName, "api"))
		apiFiles, err := getApiFiles()
		if err != nil {
			fmt.Printf("Error finding .api files: %v\n", err)
			return
		}
		runCommand(filepath.Join(projectPath, "cmd", "goctl", "goctl"), append([]string{"api", "go"}, append(apiFiles, "-dir", "./", "-style", "goZero", "-home", filepath.Join(projectPath, "common", "goctl", "1.5.0"))...)...)
	}
}

func runRpcCommand(serviceName, prodFlag string) {
	if prodFlag == "-prod" {
		runCommand("goctl", "rpc", "protoc", serviceName+".proto", "--go_out=./types", "--go-grpc_out=./types", "--zrpc_out=.", "-style", "goZero", "-home", filepath.Join(projectPath, "common", "goctl", "1.5.0"))
	} else {
		buildGoctl()
		changeDirectory(filepath.Join(projectPath, "app", serviceName, "rpc"))
		runCommand(filepath.Join(projectPath, "cmd", "goctl", "goctl"), "rpc", "protoc", serviceName+".proto", "--go_out=./types", "--go-grpc_out=./types", "--zrpc_out=.", "-style", "goZero", "-home", filepath.Join(projectPath, "common", "goctl", "1.5.0"))
	}
}

func runModelCommand(prodFlag string) {
	changeDirectory(filepath.Join(projectPath, "cmd", "updateModel"))
	if prodFlag == "-prod" {
		fmt.Println("暂无 model -prod命令")
	} else {
		runCommand("go", "run", "./", "-dev")
	}
}

func buildGoctl() {
	fmt.Println("Building goctl...")
	changeDirectory(filepath.Join(projectPath, "cmd", "goctl"))
	runCommand("go", "build", "goctl.go")
}

func changeDirectory(dir string) {
	if runtime.GOOS == "windows" {
		dir = filepath.FromSlash(dir)
	}
	fmt.Printf("Changing directory to: %s\n", dir)
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("Error changing directory to %s: %v\n", dir, err)
	} else {
		printCurrentDirectory()
	}
}

func printCurrentDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
	} else {
		fmt.Printf("Current directory: %s\n", dir)
	}
}

func runCommand(name string, args ...string) {
	fmt.Printf("Running command: %s %v\n", name, args)
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running command %s %v: %v\n", name, args, err)
	}
}

func getApiFiles() ([]string, error) {
	files, err := filepath.Glob("*.api")
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no .api files found")
	}
	return files, nil
}
