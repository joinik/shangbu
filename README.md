# go_商布

GoWeb 框架 gin-商铺项目

## 技术栈

- Go
- Gorm
- Mysql

## 项目的主要依赖：

Golang V1.20

- gin
- gorm
- mysql
- redis
- ini
- jwt-go
- crypto
- logrus
- qiniu-go-sdk
- dbresolver

## 项目计划


| 日期                | 目标                           | 完成度 |
| ------------------- | ------------------------------ | ------ |
| 2022 年 12 月 16 日 | 完成 sql 功能封装成 dao        | ✔️   |
| 2022 年 12 月 16 日 | 完成 sql 功能封装成 dao        | ✔️   |
| 2022 年 12 月 18 日 | 完成 文章功能 以及用户功能修复 | ✔️   |

## 项目结构

## 配置文件

`conf/config.ini` 文件配置

```ini
#debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000

[mysql]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = root
DbName =

[redis]
RedisDb = redis
RedisAddr = 127.0.0.1:6379
RedisPw =
RedisDbName =

[qiniu]
AccessKey =
SerectKey =
Bucket =
QiniuServer =

[email]
ValidEmail=http://localhost:8080/#/vaild/email/
SmtpHost=smtp.qq.com
SmtpEmail=
SmtpPass=
#SMTP服务的通行证

[es]
EsHost = 127.0.0.1
EsPort = 9200
EsIndex = mylog
```

## 项目运行

**本项目采用 Go Mod 管理依赖**

**下载依赖**

```go
go mod tidy
```

**项目启动**

```go
go run main.go
```
