### 创建user api
```goctl api new user```
### 编辑user.api后重新生成代码
```
#进user目录执行
goctl api go -api user.api -dir . -style goZero
#如果用了自定义模版则输入home路径
goctl api go -api user.api -dir . -style goZero -home ../../../common/goctl/1.5.0

#mac电脑可以vim 
然后加入
alias gozeroApi='goctl api go -api user.api -dir . -style goZero -home ../../../common/goctl/1.5.0'
alias gozeroModel='goctl model mysql ddl --src *.sql --dir . -style goZero'
```
### 创建user model
```
#先创建sql建表文件，如user.sql
#进入到目录下，执行model生成脚本
goctl model mysql ddl --src user.sql --dir . -style goZero
```
### 标准表模版
```
CREATE TABLE `表名`  (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
    
    `plat_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '应用id',
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
#然后根据运行命令时的目录的相对路径，用自定义模版生成uer.api
goctl api go -api user.api -dir . -style goZero -home ../../../common/goctl/1.5.0
```

## 约定
### 表字段约定

| 后缀    | 字段类型 | 数据类型     |备注 
|-------|------|----------| --------- 
| _date | 日期   | date     | 年月日
| _time | 时间   | datetime | 年月日时分秒
| _at   | 时间   | int      | 时间戳
| _num  | 数字   | int      | 不带小数点  
| _fee  | 价格   | int      | 根据系统小数点读写自动换算
| _qty  | 数量   | int      | 根据系统小数点读写自动换算 

# Todo
- 配置全局变量，如在model中，如果设置TableName()，则用表前缀模式