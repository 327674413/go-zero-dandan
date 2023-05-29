### 创建user api
```goctl api new user```
### 编辑user.api后重新生成代码
```
#进user目录执行
goctl api go -api user.api -dir . -style goZero
```
###启动user api
```
#进user目录执行
go run user.go -f etc/user-api.yaml 
```

### 获取模版
```
goctl template init
#会在指定目录生成模版文件
#然后根据运行命令时的目录的相对路径，用自定义模版生成uer.api
goctl api go -api user.api -dir . -style goZero -home ../../common/goctl/1.5.0
```
# Todo
- 配置全局变量，如在model中，如果设置TableName()，则用表前缀模式