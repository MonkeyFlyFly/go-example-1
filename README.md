# go-example-1
```

$ mkdir go-gin-example && cd go-gin-example

$ go env -w GO111MODULE=on

$ go env -w GOPROXY=https://goproxy.cn,direct

$ go mod init github.com/EDDYCJY/go-gin-example
go: creating new go.mod: module github.com/EDDYCJY/go-gin-example

$ ls
go.mod


$ go get -u github.com/gin-gonic/gin




$ go run test.go

$ go mod tidy


conf：用于存储配置文件
middleware：应用中间件
models：应用数据库模型
pkg：第三方包
routers 路由逻辑处理
runtime：应用运行时数据

```

```
go get -u github.com/go-ini/ini
go get -u github.com/unknwon/com
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql


标准库
fmt：实现了类似 C 语言 printf 和 scanf 的格式化 I/O。格式化动作（‘verb’）源自 C 语言但更简单
net/http：提供了 HTTP 客户端和服务端的实现
Gin
gin.Default()：返回 Gin 的type Engine struct{...}，里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
router.GET(…){…}：创建不同的 HTTP 方法绑定到Handlers中，也支持 POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的 Restful 方法
gin.H{…}：就是一个map[string]interface{}
gin.Context：Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
```

```
Gin：Golang 的一个微框架，性能极佳。
beego-validation：本节采用的 beego 的表单验证库，中文文档。
gorm，对开发人员友好的 ORM 框架，英文文档
com，一个小而美的工具包。

models.go

import (
	"fmt"
	"github.com/EDDYCJY/go-example-1/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
	_"github.com/go-sql-driver/mysql" //注意这里会报错需要手动引入
)

```


```
顺序
router.go --> api v1 article.go ---> models article.go

```


```
db:
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');


加入auth认证
jwt

go get -u github.com/dgrijalva/jwt-go

step：

我们在pkg下的util目录新建jwt.go
我们在middleware下新建jwt目录，新建jwt.go文件
在models下新建auth.go文件
在routers下的api目录新建auth.go文件
我们打开routers目录下的router.go

验证Token
http://127.0.0.1:8000/auth?username=test&password=test123456

{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2MzA5NTM3MzgsImlzcyI6Imdpbi1ibG9nIn0.bDimk-93IDLMC0i8oucHv-CjU6v9F2OkPEBxcfhe0WU"
    },
    "msg": "ok"
}


将中间件接入Gin
 router.go --> apiv1.Use(jwt.JWT())

验证
{
    "code": 400,
    "data": null,
    "msg": "请求参数错误"
}

127.0.0.1:8000/api/v1/stu/2?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2MzA5NTQxMzYsImlzcyI6Imdpbi1ibG9nIn0.VAOQzwRzxWPS6kIzpLioSQgwRJrj6t0KkVy0GUwqB7c
```