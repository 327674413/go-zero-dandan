Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
Log:
  Level: info
  Stat: false #关闭logx每分钟的指标统计，等价于logx.DisableStat()
Mode: dev
I18n:
  Default: zh-cn
  Langs:
    - "../../../common/lang/en_us.toml"
    - "../../../common/lang/zh_cn.toml"
