# go-project-boot

微盟云开放平台的Go项目工程启动脚手架，提供一套Go工程的标准规范，开发者更多关注业务本身，减少开发成本，使其可以快速接入微盟云的开放生态。

## 介绍

### 功能列表

* 脚手架框架
* 能力实现
  * SPI 扩展点实现、注册
  * MSG 订阅、注册
* 组件列表
  * MySQL 封装
  * Redis 封装
  * Log 组件封装
  * OAuth 封装
  * Encryption 封装
  * HttpClient 封装
  * Apollo 封装

### 项目结构

```
.
├── LICENSE.md
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── pkg
│   ├── auth		# OAuth 授权组件
│   ├── boot		# 生命周期管理框架目录
│   ├── codec		# 编解码组件，例如 JSON
│   ├── config		# 配置组件，包括本地配置与Apollo配置
│   ├── db			# 数据库组件，包括 Mysql、Redis
│   ├── encrypt		# 加解密组件
│   ├── http		# Http 组件，包括 Server HealthCheck 与 Client
│   ├── wcontext	# 服务容器，包括运行时数据库等
│   └── wlog		# 日志组件
└── test
```

## 快速开始

1. 新建空白项目，并添加依赖

	```bash
	# 创建项目并初始化
	mkdir go-demo && cd go-demo
	touch main.go
	go mod init go-demo

	# 增加配置文件
	echo server.address = :8080 > application.properties
	```

2. 添加启动入口，并实现一个简单的 Controller

	```go
	package main

	import (
		"context"
		"github.com/cloudwego/hertz/pkg/app"
		"github.com/cloudwego/hertz/pkg/app/server"
		"github.com/weimob-tech/go-project-base/pkg/http"
		"github.com/weimob-tech/go-project-boot/pkg/boot"
	)

	func main() {
		starter := boot.Starter(
			boot.WithHttpServer(),
			boot.ConfigureHttpServer(route),
		)

		starter.Start()
	}

	func route(s http.Server) {
		s.GetServer().(*server.Hertz).GET("", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(200, "Hello, world!")
		})
	}
	```

3. 本地调试启

	``` bash
	# 安装依赖
	GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod tidy
	# 启动服务
	go run ./main.go

	# 请求结果
	curl -i http://127.0.0.1
	```

4. 生产环境运行，可以在微盟云开发平台进行构建镜像，并发布到容器集群

## 使用文档

* [能力文档](http://doc.weimobcloud.com/list?tag=2396&menuId=19&childMenuId=1&isold=2)
* [开发者入驻](http://doc.weimobcloud.com/word?menuId=46&childMenuId=47&tag=2970&isold=2)
* [应用开发](http://doc.weimobcloud.com/word?menuId=53&childMenuId=54&tag=2488&isold=2)

## 贡献方法

* 申请加入weimob_tech

## 联系我们

* Weimob-tech@weimob.com
