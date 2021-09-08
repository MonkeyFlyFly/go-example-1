module github.com/EDDYCJY/go-example-1

go 1.16

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/astaxie/beego v1.12.3
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/kr/pty v1.1.5 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2 // indirect
	github.com/swaggo/gin-swagger v1.2.0 // indirect
	github.com/swaggo/swag v1.7.1 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli v1.22.5 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210907225631-ff17edfbf26d // indirect
	golang.org/x/sys v0.0.0-20210908143011-c212e7322662 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => ./go-example-1/pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => ./go-example-1/middleware
	github.com/EDDYCJY/go-gin-example/models => ./go-example-1/models
	github.com/EDDYCJY/go-gin-example/pkg/setting => ./go-example-1/pkg/setting
	github.com/EDDYCJY/go-gin-example/routers => ./go-example-1/routers
)
