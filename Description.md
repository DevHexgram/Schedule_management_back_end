# 后端使用说明


## 配置文件:


位置:`/main_code/back/src/SMS_config/conf.toml`

内容:
```toml
[DB]
Address = "localhost"           # 数据库所在的位置(ip)
User = "XXXXX"         # 数据库用户名
Password  = "XXXXX"    # 数据库用户密码
DBName = "XXXXX"          # 数据库名

[Web]
Port = ":1221"                  # 使用的接口,注意':'不能少
```


## 源代码:


位置:`./src`

模块关系:
* main 
* middleware 中间件
* pkg/e 错误码
* pkg/setting 初始化配置
* pkg/util 小工具
* routers 路由
* api api





