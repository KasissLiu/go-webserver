### GO WEBSERVER

##### 基于go语言实现的网络服务器框架

目录结构如下

```
├── config
│   ├── certs
│   │   ├── server.crt
│   │   └── server.key
│   ├── db.ini
│   └── server.ini
├── controllers
│   ├── api
│   │   ├── common.go
│   │   └── user.go
│   └── ws
│       └── serverState.go
├── dbserver
│   └── mysql.go
├── go-webserver
├── main.go
├── models
│   ├── servers.go
│   └── user.go
├── readme.md
├── routes
│   ├── api.go
│   ├── web.go
│   └── ws.go
├── servers
│   ├── dynamicServer.go
│   ├── fileServer.go
│   ├── httpServer.go
│   └── initServer.go
└── www
    ├── favicon.ico
    └── index.html

```

实现功能

- 静态文件服务器
- 业务逻辑处理器
- websocket服务器
- 配置文件载入
- 数据库链接池

