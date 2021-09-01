module github.com/EDDYCJY/go-example-1

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.62.0 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => ./go-example-1/pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => ./go-example-1/middleware
	github.com/EDDYCJY/go-gin-example/models => ./go-example-1/models
	github.com/EDDYCJY/go-gin-example/pkg/setting => ./go-example-1/pkg/setting
	github.com/EDDYCJY/go-gin-example/routers => ./go-example-1/routers
)
