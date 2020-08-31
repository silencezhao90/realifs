# realifs

## Introduction
对各种存储端（阿里云oss、百度云、腾讯云、华为云、minio等）进行封装。根据不同配置调用对应不同端存储功能实现文件管理。

接入端存储端
- [X] 阿里云oss
- [X] minio
- [ ] 腾讯云oss
- [ ] 百度云oss
- [ ] AWS oss
- [ ] 华为云oss

## Quick start
### 前置条件
* 确保云端存储服务已开通或本地minio已经设置启动
* 将存储配置写入配置文件

### 配置文件说明
配置文件位于当前目录下config.yaml

**配置文件结构**
```yaml
aliyun:                 // 阿里云配置
  accessKeyID: ""       // 填写aliyun accessKey
  accessKeySecret: ""   // 填写aliyun accessKeySecret
  bucketName: ""        // oss bucket name
  externalEndpoint: ""  // oss 对外endpoint
  internalEndpoint: ""  // oss 内网endpoint

minio:                  // minio配置
  endpoint: ""          // 对象存储服务的URL
  accessKeyID: ""       // Access key是唯一标识你的账户的用户ID。
  secretAccessKey: ""   // Secret key是你账户的密码。
  secure: false         // true代表使用HTTPS
  bucketName: ""

default:
  storage: aliyun   // 默认使用的存储端
```

### 运行程序
默认端口8080
```go
go build -o server main.go
```
```go
./server
```
** 可指定配置文件和端口运行 **
```go
./server --port=8080 --config=./config.yaml
```

### 生成接口文档
```go
swag init -g router/router.go -o docs
```

## Doc
服务运行后可访问接口文档http://localhost:8080/swagger/index.html
