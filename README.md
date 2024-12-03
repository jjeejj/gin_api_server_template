# gin_api_server_template
golang gin api server template

## 项目结构

```
.
├── LICENSE
├── Makefile # make 命令, 相关的脚本
├── README.md
├── app # 项目代码, 业务逻辑
│   ├── api # api 层，对外接口实现逻辑，按照模块划分
│   │   ├── app.go
│   │   ├── init.go
│   │   └── user.go
│   ├── app.go # 项目启动入库
│   ├── docs # swagger 文档
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── event # 事件
│   │   ├── event
│   │   └── listener
│   ├── middleware # 中间件
│   │   ├── authorization.go
│   │   ├── cors.go
│   │   └── request_log.go
│   ├── model # 数据库 模型定义
│   │   ├── base.go
│   │   └── user
│   │       └── user.go
│   ├── repository # 数据库操作 repo 层
│   │   └── user
│   │       └── user.go
│   ├── request # 请求参数
│   │   ├── app.go
│   │   ├── base.go
│   │   └── user
│   │       └── user.go
│   ├── response # 响应参数以及相关封装
│   │   ├── app.go
│   │   ├── errcode.go
│   │   ├── response.go
│   │   └── user
│   │       └── user.go
│   ├── route # 路由定义
│   │   ├── app.go
│   │   ├── route.go
│   │   └── user.go
│   ├── service # 对应模块公共的逻辑，抽离出来复用模块
│   │   └── user.go
│   ├── struct # 封装一些公共的结构体
│   │   ├── page.go
│   │   └── user_token.go
│   └── task # 定时任务
│       └── task.go
├── cmd
│   └── run.go
├── config.yaml # 配置文件
├── config.yaml.example # 配置文件示例
├── go.mod
├── go.sum
├── internal # 内部包
│   ├── bootstrap # 启动初始化
│   │   ├── init.go
│   │   └── migrate.go
│   ├── config # 配置结构体定义以及初始化
│   │   └── config.go
│   ├── const # 全局常量
│   │   ├── const.go
│   │   └── context.go
│   ├── global # 全局变量
│   │   └── global.go
│   ├── logger # 日志封装
│   │   └── logger.go
│   ├── mysql # mysql 封装
│   │   └── mysql.go
│   ├── redis # redis 封装
│   │   └── redis.go
│   └── server # http server
│       ├── http.go
│       └── route.go
├── main.go # 入口文件
├── pkg # 工具包
└── test # 测试代码
```

## 开发流程
