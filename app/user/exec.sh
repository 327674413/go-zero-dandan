rootPath="/Users/yelin/go_dev/project/src/go-zero-dandan"
appPath="${rootPath}/app/user/rpc"
goctlPath="${rootPath}/common/goctl/1.5.0"
goctl rpc protoc user.proto --go_out=${appPath}/types --go-grpc_out=${appPath}/types --zrpc_out=${appPath} -style goZero -home ${goctlPath}

# 直接cmd跑用下面
goctl rpc protoc ./app/user/rpc/user.proto --go_out=./app/user/rpc/types --go-grpc_out=./app/user/rpc/types --zrpc_out=./app/user/rpc -style goZero -home ./common/goctl/1.5.0
