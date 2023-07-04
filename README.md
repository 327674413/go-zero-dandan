### 创建user api
```goctl api new user```
### 编辑user.api后重新生成代码
```
#进user目录执行
goctl api go -api user.api -dir . -style goZero
#如果用了自定义模版则输入home路径，在desc使用
goctl api go -api user.api -dir ../ -style goZero -home ../../../../common/goctl/1.5.0

#mac电脑可以vim ~/.bash_profile 
然后加入
alias gozeroApi='goctl api go -api *.api -dir ../ -style goZero -home ../../../common/goctl/1.5.0'
```
### 创建user model
```
#先创建sql建表文件，如user.sql
#进入到目录下，执行model生成脚本
goctl model mysql ddl --src user.sql --dir . -style goZero -home ../../../common/goctl/1.5.0
#或者直接连接数据库创建
goctl model mysql datasource --ignore-columns="delete_at" -url="${DB_USER}:${$DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -table="${DB_TABLE}" . -style goZero -home ../../../common/goctl/1.5.0
```

### rpc服务创建
```
#在rpc目录里新建一个proto文件，然后执行
goctl rpc protoc message.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. -style goZero -home ../../../common/goctl/1.5.0

```
### 标准表模版
```
CREATE TABLE `表名`  (
    `id` bigint UNSIGNED NOT NULL,
    
    `plat_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '应用id',
    `create_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间戳',
    `update_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间戳',
    `delete_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间戳',
    PRIMARY KEY (`id`)
);
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
#然后根据运行命令时的目录的相对路径，用自定义模版生成uer.api(用*就会该目录下所有.api文件)
goctl api go -api *.api -dir . -style goZero -home ../../../common/goctl/1.5.0
```

## 约定
### 表字段约定

| 后缀    | 字段类型 | 数据类型     | 备注            |
|-------|------|----------|---------------|
| _date | 日期   | date     | 年月日           |
| _time | 时间   | datetime | 年月日时分秒        |
| _at   | 时间   | int      | 时间戳           |
| _num  | 数字   | int      | 不带小数点         |
| _fee  | 价格   | int      | 根据系统小数点读写自动换算 |
| _qty  | 数量   | int      | 根据系统小数点读写自动换算 |

# 各中间件工具查看
- Prometheus监控：http://localhost:9090/targets?search=

# Todo
- common/resd汇总的msg.go拆分成代码生成和手工两部分，对于错误提示可以用代码生成批处理
- 接管httpx.Parse，解决不能同时支持数字和字符串数字的问题以及转化失败报错英文的问题
- 研究一下APISIX  和 kong ，实现统一的网关入口，做黑名单等事务
- 如何获取框架自带的tracer自行使用
- kafka在docker里必须要添加hosts把容器id放进去，不然找不到地址的问题解决
- 研究一下gozero的gateway，直接转发rpc，不需要api再定义一遍，不然很多业务，api接口一次结构体，rpc定义一次结构体太难受了
- 研究goctl，解决：
- - 部分表是不需要plat_id的，现在生成都会有，需要能自动化处理该问题，现在分成两套模版生成，有点难控制
- - 一旦模版修改，需要每个文件都要去跑一遍的问题