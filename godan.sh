#!/bin/bash

# 公共路径变量
PROJECT_PATH="/Users/yelin/go_dev/project/src/go-zero-dandan"

# 判断命令类型并设置服务名称
if [ "$1" == "model" ]; then
    SERVICE_NAME=""
else
    if [ -z "$2" ]; then
        echo "Usage: $0 {api service_name [-prod] | rpc service_name [-prod] | model [-prod]}"
        exit 1
    fi
    SERVICE_NAME="$2"
fi

# 根据第一个参数执行相应的命令
case $1 in
    api)
        if [ "$3" == "-prod" ]; then
            # 直接执行命令
            goctl api go -api *.api -dir ./ -style goZero -home ../../../common/goctl/1.5.0
        else
            # 开发模式构建并执行命令
            cd $PROJECT_PATH/cmd/goctl/
            go build goctl.go
            cd $PROJECT_PATH/app/$SERVICE_NAME/api
            $PROJECT_PATH/cmd/goctl/goctl api go -api *.api -dir ./ -style goZero -home ../../../common/goctl/1.5.0
        fi
        ;;
    rpc)
        if [ "$3" == "-prod" ]; then
           # 直接执行命令
           goctl rpc protoc $SERVICE_NAME.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. -style goZero -home ../../../common/goctl/1.5.0
        else
            # 开发模式构建并执行命令
            cd $PROJECT_PATH/cmd/goctl/
            go build goctl.go
            cd $PROJECT_PATH/app/$SERVICE_NAME/rpc
            $PROJECT_PATH/cmd/goctl/goctl rpc protoc $SERVICE_NAME.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. -style goZero -home ../../../common/goctl/1.5.0
        fi
        ;;
    model)
        cd $PROJECT_PATH/cmd/updateModel/
        if [ "$2" == "-prod" ]; then
            echo "暂无 model -prod命令"
        else
            go run ./ -dev
        fi
        ;;
    *)
        echo "Usage: $0 {api service_name [-prod] | rpc service_name [-prod] | model [-prod]}"
        exit 1
        ;;
esac