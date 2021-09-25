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


```
日志
logging --> log.go --> file.go 

业务调用
"github.com/EDDYCJY/go-example-1/pkg/logging"

logging.Info(data, "------------------")
```


```
swagger: 
安装 swag
go get -u github.com/swaggo/swag/cmd/swag@v1.6.5

swag -v


安装 gin-swagger
go get -u github.com/swaggo/gin-swagger@v1.2.0 
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template

swag init

```

#### 定制 GORM Callbacks
```
GORM 本身是由回调驱动的，所以我们可以根据需要完全定制 GORM，以此达到我们的目的

注册一个新的回调
删除现有的回调
替换现有的回调
注册回调的顺序

在 GORM 中包含以上四类 Callbacks，我们结合项目选用 “替换现有的回调” 来解决一个小痛点

CreatedOn : 时间
models.go


```


#### Cron
```

Cron 表达式格式
字段名	是否必填	允许的值	允许的特殊字符
秒（Seconds）	Yes	0-59	* / , -
分（Minutes）	Yes	0-59	* / , -
时（Hours）	Yes	0-23	* / , -
一个月中的某天（Day of month）	Yes	1-31	* / , - ?
月（Month）	Yes	1-12 or JAN-DEC	* / , -
星期几（Day of week）	Yes	0-6 or SUN-SAT	* / , - ?

```

```
Cron 特殊字符
1、星号 ( * )

星号表示将匹配字段的所有值

2、斜线 ( / )

斜线用户 描述范围的增量，表现为 “N-MAX/x”，first-last/x 的形式，例如 3-59/15 表示此时的第三分钟和此后的每 15 分钟，到 59 分钟为止。即从 N 开始，使用增量直到该特定范围结束。它不会重复

3、逗号 ( , )

逗号用于分隔列表中的项目。例如，在 Day of week 使用“MON，WED，FRI”将意味着星期一，星期三和星期五

4、连字符 ( - )

连字符用于定义范围。例如，9 - 17 表示从上午 9 点到下午 5 点的每个小时

5、问号 ( ? )

不指定值，用于代替 “ * ”，类似 “ _ ” 的存在，不难理解
```

```
预定义的 Cron 时间表
输入	简述	相当于
@yearly (or @annually)	1 月 1 日午夜运行一次	0 0 0 1 1 *
@monthly	每个月的午夜，每个月的第一个月运行一次	0 0 0 1 * *
@weekly	每周一次，周日午夜运行一次	0 0 0 * * 0
@daily (or @midnight)	每天午夜运行一次	0 0 0 * * *
@hourly	每小时运行一次	0 0 * * * *
```

#### 利用cron进行硬删除
```
注意硬删除要使用 Unscoped()，这是 GORM 的约定

func CleanAllStu() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Stu{})

	return true
}

time.NewTimer + for + select + t1.Reset
如果你是初学者，大概会有疑问，这是干嘛用的？

**（1）time.NewTimer **

会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息

（2）for + select

阻塞 select 等待 channel

（3）t1.Reset

会重置定时器，让它重新开始计时

```


#### 第一步
```
将散落在其他文件里的配置都删掉，统一在 setting 中处理以及修改 init 函数为 Setup 方法

打开 pkg/setting/setting.go 文件

将 models.go、setting.go、pkg/logging/log.go 的 init 函数修改为 Setup 方法
将 models/models.go 独立读取的 DB 配置项删除，改为统一读取 setting
将 pkg/logging/file 独立的 LOG 配置项删除，改为统一读取 setting
```

#### 第二步
```
在这一步我们要设置初始化的流程，打开 main.go 文件，修改内容
```