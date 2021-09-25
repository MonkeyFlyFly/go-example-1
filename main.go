package main

import (
	"github.com/EDDYCJY/go-example-1/models"
	"github.com/EDDYCJY/go-example-1/pkg/logging"
	"github.com/EDDYCJY/go-example-1/pkg/setting"
)

func main() {
	//router := routers.InitRouter()

	/***
	Addr：监听的 TCP 地址，格式为:8000
	Handler：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
	TLSConfig：安全传输层协议（TLS）的配置
	ReadTimeout：允许读取的最大时间
	ReadHeaderTimeout：允许读取请求头的最大时间
	WriteTimeout：允许写入的最大时间
	IdleTimeout：等待的最大时间
	MaxHeaderBytes：请求头的最大字节数
	ConnState：指定一个可选的回调函数，当客户端连接发生变化时调用
	ErrorLog：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为nil则默认以日志包的标准日志记录器完成（也就是在控制台输出）
	*/
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//s.ListenAndServe()

	setting.Setup()
	models.Setup()
	logging.Setup()
}
