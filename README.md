# gin_api_server_template
golang gin api server template

## 项目结构

```
├── LICENSE
├── Makefile
├── README.md
├── app
│   ├── api
│   │   ├── app.go
│   │   ├── init.go
│   │   └── user.go
│   ├── app.go
│   ├── docs
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── event
│   │   ├── event
│   │   └── listener
│   ├── middleware
│   │   ├── authorization.go
│   │   ├── cors.go
│   │   └── request_log.go
│   ├── model
│   │   ├── base.go
│   │   └── user
│   │       └── user.go
│   ├── repository
│   │   └── user
│   │       └── user.go
│   ├── request
│   │   ├── app.go
│   │   ├── base.go
│   │   └── user
│   │       └── user.go
│   ├── response
│   │   ├── app.go
│   │   ├── errcode.go
│   │   ├── response.go
│   │   └── user
│   │       └── user.go
│   ├── route
│   │   ├── app.go
│   │   ├── route.go
│   │   └── user.go
│   ├── service
│   │   └── user.go
│   └── struct
│       ├── page.go
│       └── user_token.go
├── cmd
│   └── run.go
├── config.yaml
├── config.yaml.example
├── go.mod
├── go.sum
├── internal
│   ├── bootstrap
│   │   ├── init.go
│   │   └── migrate.go
│   ├── config
│   │   └── config.go
│   ├── const
│   │   ├── const.go
│   │   └── context.go
│   ├── global
│   │   └── global.go
│   ├── logger
│   │   └── logger.go
│   ├── mysql
│   │   └── mysql.go
│   ├── redis
│   │   └── redis.go
│   └── server
│       ├── http.go
│       └── route.go
├── main.go
├── pkg
├── test
```

