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



## 2 Directory Structure

```
.
├── conf
│   └── dev.yaml
├── dev
│   └── db.md
├── logs
│   └── info.log.201907230000
├── main.go
├── models
│   ├── demo.go
│   ├── demo_test.go
│   └── model.go
├── pkgs
│   ├── file
│   │   └── file.go
│   ├── logging
│   │   └── logging.go
│   └── setting
│       └── setting.go
├── README.md
├── routers
│   ├── api
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
- [X] API测试
- [ ] GET POST PUT DELETE 请求
- [ ] JSON FORM PARAM 三种请求方式
- [ ] 中间件[交易时间统计]
- [ ] Redis 数据
- [ ] 容器化部署
- [ ] 日志集中收集 ELK
- [X] 安全机制 文件上传、下载安全限制
- [ ] AUTH2
- [ ] 服务注册 

## 5 Stargazers


[![Stargazers over time](https://starchart.cc/spiolynn/go-web-scaffold.svg)](https://starchart.cc/spiolynn/go-web-scaffold)


## 6 demo api

> https://documenter.getpostman.com/view/6107364/SVSRFR9z?version=latest

