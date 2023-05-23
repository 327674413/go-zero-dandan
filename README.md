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
