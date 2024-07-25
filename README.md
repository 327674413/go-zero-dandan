这是一个使用go-zero微服务框架的尝试项目，需要有go-zero的基础。
# 项目目标
- 像写单体服务一样开发微服务
- 代码生成业务
- 集成各类业务的基础组件

# 业务设想
- 用户服务：所有业务的底层服务，作为所有服务的依赖，共用一套用户体系（虽然微服务要解耦，但一个人写还根据业务模块建立不同的用户体系太伤了）
- 权限服务：搭配用户服务的权限配置，功能模块权限、操作行为权限、字段权限、数据权限
- IM服务：提供最基本的消息通讯能力，用于实现基本的客服系统功能、业务消息推送通知等即时通讯
- B2C商城服务：常规的b端商城，基本的商品、订单、支付、活动、积分、优惠券等
- 公式解析：用于定义公式进行判断或计算，目前看到goja，作用在流程引擎中支持定义公式来决定流程分支
- 工作流引擎：接入或自制常用工作流，用于后台管理系统的审批
- 灰度发布：实现灰度发布机制
- 业务模版：面向多变性的C端业务的代码生成，提供基本的参数校验、简单业务代码生成
- 业务解析：面向规律性的管理端或数据记录类管理系统的代码解析，通过约定json，实现crud功能，可放置在本地写前后端业务、可放置配置中心方便改业务、可放置放到数据库配合可视化编辑器在线修改业务
- 其他…………
# 已研究部分
- 集成i18n，支持国际化，通过维护主语言包，自动生成和更新
- 改造goctl工具与模版，生成自己专用的业务代码
- 自制orm，用thinkphp model的方式操作数据库
- 参照api语法，自制rpc文件，生成rpc的自动化业务代码

# 近期代办
- ws前端离线，发送或心跳前主动重连或提示用户主动重连
- 如果接受好友申请，通知对方刷新好友列表
- ws里也接入国际化 与 打印调试
- 确认下api和rpc 用户信息那块是否需要init
- 完善im基础功能：好友、群组、聊天、公告、已读状态
- 对rpc没启动时，报错信息是produced zero addresses的封装
- trace_id需要延续到kafka吗？如何弄？
# 笔记

## 代码生成
自己重新改了goctl工具，在cmd/goctl目录下，模版文件在common/goctl/1.5.0下，可以将项目目录下的godan.sh放到gobin目录里，然后建立软链接：

```
sudo ln -s /Users/yelin/go_dev/project/src/go-zero-dandan/godan.sh /usr/local/bin/godan
```
建立完毕后，可以直接通过godan命令生成代码(如果有重名连接，可以sudo rm /usr/local/bin/godan 来删除)
```
#生成user-api ，如果是改了模版全量刷新，就更新godan.sh下的apiList,然后执行godan api all
godan api user

#生成user-rpc ，如果是改了模版全量刷新，就更新godan.sh下的rpcList,然后执行godan rpc all
godan rpc user

#生成model
godan model

#修改common/lang中的语言文件后，更新语言包
godan lang
```


## goctl二次开发

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

# 各中间件工具查看
## Prometheus监控
地址：http://localhost:9090/targets?search=
## minio
web管理地址 http://localhost:9001/
密码要求8位，docker-compose中有设置，root 12345678
## 查看etcd服务情况
先进容器：docker exec -it 容器id /bin/bash
然后执行：etcdctl get --prefix "" --keys-only=true
## 配置中心sail
地址：http://localhost:8108/ui/login  admin  dandan
## 网关apisix
地址：http://localhost:9002    admin  admin
### MQ中间件Kafka
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

## 开发说明
### 接口相关
- api的请求入参，目前是用框架的httpx解析，所以参数可选得用optional
- api的返回值，如果想让返回的内容为nil时不会返回，则用omitempty
- 对于接口中，非必填的字段，都建议用指针类型，既明确是可选参数，又可以判断前端是否有传

### 数据库相关
- 单条数据查询，未查到的err也是nil，只有查询异常err才有内容。如果要判断是否有查到数据，就判断数据是不是nil。
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



## 约定
### 接口约定
- getOne 查询单个信息
- getPage 查询列表，分页
- getList 查询列表，所有
### 数据库表结构
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
### 表字段约定

| 后缀    | 字段类型 | 数据类型     | 备注            |
|-------|------|----------|---------------|
| _date | 日期   | date     | 年月日           |
| _time | 时间   | datetime | 年月日时分秒        |
| _at   | 时间   | int      | 时间戳           |
| _num  | 数字   | int      | 不带小数点         |
| _fee  | 价格   | int      | 根据系统小数点读写自动换算 |
| _qty  | 数量   | int      | 根据系统小数点读写自动换算 |

### 国际化约定
- 模版支持参数二次解析的，但需要传入的是Var开头的变量，如果特殊变量本身是Var开头的单词却不想二次转意，则用*Var来写，会自动去掉*



# 待研究部分
- gozero的kq发送消息是异步的，无法知道是否推送成功，有些消息队列场景需要确保投递成功的
- 走通采集logx日志到es，根据traceid，实现搜索某个用户的所有操作 和 某张表数据的变化
- 如果h5请求添加了自定义头，还是会跨域，只能用官方默认支持的AccessToken, Token 作为自定义头
- goland配置远程自动上传，删除文件时好像不会触发
- kafka在docker里必须要添加hosts把容器id放进去，不然找不到地址的问题解决
- 用goland终端启动rpc服务，放一段时间不管，比如笔记本盖上后，回来打开，服务没退出，但连不上rpc，需要人工重启。（etcd里的服务、普罗米修斯里的服务都是正常的）
