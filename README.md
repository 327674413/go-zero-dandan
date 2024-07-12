# 项目目标
- 像写单体服务一样开发微服务
- 集成各类业务的基础组件

# 已完成部分
- 集成i18n，支持国际化
- 改造goctl，生成自己专用的业务代码
- 自制orm，用thinkphp model的方式操作数据库
- 参照api语法，自制rpc文件，生成rpc的自动化业务代码

# 近期待办
- 完善resd，i18n返回中文好像又有问题，rpc返回，并且支持Info和Err， 正常业务校验可以Info类型，确实报错用err
- gozero的kq发送消息是异步的，无法知道是否推送成功，有些消息队列场景需要确保投递成功的
- 走通采集logx日志到es，根据traceid，实现搜索某个用户的所有操作 和 某张表数据的变化
# 笔记

## 代码生成
自己重新改了goctl工具，在cmd/goctl目录下，模版文件在common/goctl/1.5.0下，可以将项目目录下的godan.sh放到gobin目录里，然后建立软链接：

```
sudo ln -s /Users/yelin/go_dev/project/src/go-zero-dandan/godan.sh /usr/local/bin/godan
```
建立完毕后，可以直接通过godan命令生成代码,如果有重名连接，可以sudo rm /usr/local/bin/godan 来删除
```
#生成user-api
godan api user

#生成user-rpc
godan rpc user

#生成model
godan model
```

### 标准表模版
```
CREATE TABLE `表名`  (
    `id` char(20) NOT NULL,
    `remark` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '备注',
    `plat_id` char(20) NOT NULL DEFAULT '' COMMENT '应用id',
    `create_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间戳',
    `update_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间戳',
    `delete_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间戳',
    PRIMARY KEY (`id`)
);

```
### goctl二次开发

```
# 以修改model生成规则中，将bigint类型的id 变成string类型为例子
# 首先git上下载goct源码，以1.5.0版本为例子
# 进入goctl目录中的model/sql/parser/parser.go
# 在大概360行附近，添加：
if each.Name == "id" {
    dt = "string"
}
# 这里的each就是字段，each.Name是数据库的字段名，dt是经过gozero原先的映射关系处理后得到的go数据类型，这里强制变成string
# 改完后，编译文件，然后通过which看一下goctl目录放哪，将编译后的文件覆盖，就可以用自己二次开发的goctl了

# 调试的话，用 go run goctl.go 后面一样可以跟上原先的参数
# 但里面用好像只能用fmt.Println()才能看空值台打印，用logx.Info好像不行

```

### 启动user api
```
#进user目录执行
go run user.go -f etc/user-api.yaml 
```
### docker-compose启动其他名字配置
```
#小写-d好像不行
docker-compose -f docker-compose-test.yml up --detach
```

### kafka
```
#删除topic
docker exec -it kafka kafka-topics.sh --delete --topic 主题名 --bootstrap-server 127.0.0.1:9092
#创建topic
docker exec -it kafka kafka-topics.sh --create --topic 主题名 --bootstrap-server 127.0.0.1:9092 --partitions 3 --replication-factor 1
#往docker里的kafka测试发消息
docker exec -it kafka kafka-console-producer.sh --broker-list 127.0.0.1:9092 --topic 主题名
#查看当前所有topic
docker exec -it kafka kafka-topics.sh --list --bootstrap-server 127.0.0.1:9092
```
### 获取模版
```
goctl template init
#会在指定目录生成模版文件
#然后根据运行命令时的目录的相对路径，用自定义模版生成uer.api(用*就会该目录下所有.api文件)
goctl api go -api *.api -dir . -style goZero -home ../../../common/goctl/1.5.0
```
### 网关apisix说明
```
地址：http://localhost:9002    admin  admin

```
### 配置中心sail
```
地址：http://localhost:8108/ui/login  admin  dandan
```

## 开发说明
### 接口相关
- api的请求入参，目前是用框架的httpx解析，所以参数可选得用optional
- api的返回值，如果想让返回的内容为nil时不会返回，则用omitempty
- 对于接口中，非必填的字段，都建议用指针类型，既明确是可选参数，又可以判断前端是否有传

### 数据库相关
- 单条数据查询，未查到的err也是nil，按照数据非nil则正常，如果是nil可以直接笼统报错未找到数据，但如果未找到需要新增时，则额外判断err是不是nil，是nil就是没数据可新增，非nil则是sql操作报错，异常报错
### 缓存相关
- 数据未查到不报错，需要判断查出来的东西是不是空字符串或者写入的目标结构体是不是有值来判断，暂时好像没碰到要区分 未查到的这种场景
## 部署说明

### 普通应用部署
```
#在目录下执行：
GOOS=linux GOARCH=amd64 go build -o fileName
```

### Docker部署
```
# 1先构建dockerfile，在服务目录运行，得到Dockerfile文件
goctl docker -go 服务文件.go  

# 2编辑生成的Dockerfile 做几处修改
（1）修改顶部为FROM golang:1.20.5-alpine AS builder 因为项目用了泛型，版本要高些，另外RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories的目录得用这个版镜像才有
（2）在COPY地方增加一行：COPY common/land /common/land  因为要把语言包文件复制进去，不然加载不到
（3）在下面的COPY也要加一行：COPY --from=builder /common/land /common/land 因为是二次构建的

# 3在项目根目录，创建docker-compose文件，service下添加配置
  asset-api:
    build:
      dockerfile: ./app/asset/api/Dockerfile
    environment:
      - TZ=Asia/Shanghai
    privileged: true
    ports:
      - "8803:8803"
    stdin_open: true
    tty: true
    networks:
      - dandan_net
    restart: always
  
#4 然后运行，如果文件名是别名，则用
docker-compose -f docker-compose-test.yml up --detach

```

## minio
web管理地址 http://localhost:9001/
密码要求8位，docker-compose中有设置，root 12345678

## 约定
## 接口约定
- getOne 查询单个信息
- getPage 查询列表，分页
- getList 查询列表，所有
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

# 查看etcd服务
先进容器：docker exec -it 容器id /bin/bash
然后执行：etcdctl get --prefix "" --keys-only=true

# Todo
## 功能todo
- 增加针对ip或ip+agent，用内存或redis，定义按秒、分或n分内，对api超过多少次请求就限流阻断的功能
- 搞清楚apisix，到底网关转发，对性能影响，是否所有请求都经过apisix转发到api再到rpc？还是可以直接由apisix转发到rpc？还是可以写一个go插件，在apisix地方就转掉？
## 问题优化研究todo
- 如果h5请求添加了自定义头，还是会跨域，只能用官方默认支持的AccessToken, Token 作为自定义头
- goland配置远程自动上传，删除文件时好像不会触发
- common/resd汇总的msg.go拆分成代码生成和手工两部分，对于错误提示可以用代码生成批处理
- 接管httpx.Parse，解决不能同时支持数字和字符串数字的问题以及转化失败报错英文的问题
- 如何获取框架自带的tracer自行使用
- kafka在docker里必须要添加hosts把容器id放进去，不然找不到地址的问题解决
- resd优化，多语言如何返回提示，如果内部有封装，内部就调不到lang，直接返回一个err外部无法识别具体报错
- log在哪写？如果在外面写，提示的行就是外面的，如果里面有很多地方有err，根本不知道哪个地方返回的。但如果写里面，那么ctx传递又是大问题，然后嵌套调用可能又会重复记录
- 用goland终端启动rpc服务，放一段时间不管，比如笔记本盖上后，回来打开，服务没退出，但连不上rpc，需要人工重启。（etcd里的服务、普罗米修斯里的服务都是正常的）
- 研究goctl，解决：
- - 部分表是不需要plat_id的，现在生成都会有，需要能自动化处理该问题，现在分成两套模版生成，有点难控制
- - 一旦模版修改，需要每个文件都要去跑一遍的问题
