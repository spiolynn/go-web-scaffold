[![github](https://img.shields.io/badge/github-spiolynn-brightgreen.svg)](http://panzi.online)

# go-web应用脚手架


## 0 quick start

```
- step1: 拉代码 需要在 $GOPATH/src 目录下
git clone https://github.com/spiolynn/go-web-scaffold.git

- step2: 下载依赖
govendor sync
(如无govendor命令,执行: go get -u github.com/kardianos/govendor)

- step3: build
go build main.go

- step4: run
./main

```


## 1 Features


### 1 依赖

```
github.com/Pallinder/go-randomdata
github.com/astaxie/beego/validation
github.com/dgrijalva/jwt-go
github.com/gin-contrib/sse
github.com/gin-gonic/gin
github.com/gin-gonic/gin/binding
github.com/gin-gonic/gin/internal/json
github.com/gin-gonic/gin/render
github.com/go-ini/ini
github.com/go-sql-driver/mysql
github.com/goinggo/mapstructure
github.com/golang/protobuf/proto
github.com/jinzhu/gorm
github.com/jinzhu/gorm/dialects/mysql
github.com/jinzhu/inflection
github.com/keepeye/logrus-filename
github.com/konsorten/go-windows-terminal-sequences
github.com/lestrrat-go/file-rotatelogs
github.com/lestrrat-go/file-rotatelogs/internal/option
github.com/lestrrat-go/strftime
github.com/mattn/go-isatty
github.com/pkg/errors
github.com/rifflock/lfshook
github.com/sirupsen/logrus
github.com/t-tomalak/logrus-easy-formatter
github.com/ugorji/go/codec
go-web-frame/pkgs/setting
go.uber.org/atomic
go.uber.org/multierr
go.uber.org/zap
go.uber.org/zap/buffer
go.uber.org/zap/internal/bufferpool
go.uber.org/zap/internal/color
go.uber.org/zap/internal/exit
go.uber.org/zap/zapcore
golang.org/x/sys/unix
gopkg.in/ffmt.v1
gopkg.in/go-playground/validator.v8
gopkg.in/yaml.v2
```

### 2 主要目录

- conf: 程序配置(yaml格式)
- dev : 开发者文档
- logs: 日志
- middlewares: 自定义中间件
- models: 模型层
- pkgs:  公共功能
- routes: 路由层
- runtime: 服务在线生成
- service: 服务实现层



## 2 Directory Structure

```
.
├── build.sh
├── conf
│   ├── dev.yaml
├── dev
│   ├── db.md
│   ├── go-web-scaffold.postman_collection.json
│   └── Restful.md
├── logs
│   ├── info.log -> logs/info.log.201907230000
│   └── info.log.201907230000
├── main.go
├── middlewares
│   └── jwt
│       └── jwt.go
├── models
│   ├── demo.go
│   ├── demo_test.go
│   └── model.go
├── pkgs
│   ├── app
│   │   └── response.go
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── file
│   │   └── file.go
│   ├── logging
│   │   └── logging.go
│   ├── security
│   │   └── jwt.go
│   │   └── jwt_test.go
│   ├── setting
│   │   └── setting.go
│   ├── upload
│   │   ├── upload.go
│   │   └── upload_test.go
│   └── util
│       └── md5.go
├── README.md
├── routers
│   ├── api
│   │   ├── auth.go
│   │   ├── upload.go
│   │   └── v1
│   │       └── control.go
│   └── route.go
└── vendor
    └── vendor.json


```

## 3 How to Use



## 4 TO-DO List

- [x] yaml配置文件
- [x] 日志模块
- [x] 日志日期分割、日志级别分割
- [x] 数据库CRUD (mysql)
- [X] TDD: 单元测试
- [X] 性能测试
- [X] API测试 PostMan-connection-json
- [X] JWT 认证
- [X] GET POST PUT DELETE 请求
- [X] JSON FORM PARAM 三种请求方式
- [X] 中间件[交易时间统计]
- [ ] Redis 数据
- [ ] 容器化部署
- [ ] 日志集中收集 ELK
- [X] 安全机制 文件上传、下载安全限制
- [ ] AUTH2
- [ ] 服务注册
- [x] 数据结构映射
    - [X] json -> struct
    - [X] struct -> json
    - [X] map -> struct
    - [X] strcut -> map
    - [X] map -> json
    - [X] json -> map

## 5 Stargazers


[![Stargazers over time](https://starchart.cc/spiolynn/go-web-scaffold.svg)](https://starchart.cc/spiolynn/go-web-scaffold)


## 6 demo api

> https://documenter.getpostman.com/view/6107364/SVSRFR9z?version=latest

