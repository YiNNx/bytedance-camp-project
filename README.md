
# Waterlang-Tiktok

2023 Btyedance Camp Group Project


```
├── cmd                    // 业务代码
│   ├── api                // api server, 负责处理http请求，转发到微服务端				
│   │   ├── controller     // Gin的控制器函数，处理http请求与响应
│   │   ├── rpc            // 作为client端调用rpc函数
│   │   ├── router         // Gin定义请求路由
│   │   └── main.go        // 启动api server直接go run此文件即可（工作目录为项目根目录
│   ├── user               // 微服务端 - 用户系统
│   │   ├── build.sh       // kitex自动生成，用于运行该用户系统服务
│   │   ├── command        // 用户系统gorm相关代码
│   │   ├── handler.go     // 放有被api server调用的rpc函数，这些函数实现主要业务逻辑
│   │   ├── kitex.yaml     // kitex自动生成
│   │   ├── main.go        // kitex自动生成，作为kitex server端与client端（api server）通信
│   │   └── script         // kitex自动生成
│   ├── comment            // 同用户系统
│   ├── favorite
│   └── video
├── config                 // 配置文件
├── db                     // 数据库
│   ├── user.go            // 定义结构体用于建表和数据操作，下同
│   ├── comment.go
│   ├── context.go
│   ├── favorite.go
│   ├── postgres.go        // 连接数据库 & 事务处理
│   └── video.go
├── idl                    // thrift接口定义语言，用于定义rpc函数接口
│                             从而使用kitex生成rpc手脚架代码
│   ├── comment.thrift
│   ├── favorite.thrift
│   ├── user.thrift      
│   └── video.thrift
├── kitex_gen              // kitex自动生成
├── middleware             // Gin相关中间件函数 
│   └── jwt.go
├── pkg                    // 相关组件代码
│   ├── bcrypt             // 哈希加密密码
│   ├── config             // 读取配置文件
│   ├── jwt                // 生成用户token
│   ├── log                // 日志系统
│   └── response           // 格式化http响应
├── docker-compose.yml     // docker-compose文件
├── LICENSE
├── README.md
└── run.sh                 // 配置了docker的linux系统可运行该sh文件一键启动项目

```

