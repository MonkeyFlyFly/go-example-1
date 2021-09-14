package routers

import (
	_ "github.com/EDDYCJY/go-example-1/docs"
	"github.com/EDDYCJY/go-example-1/middleware/jwt"
	"github.com/EDDYCJY/go-example-1/pkg/setting"
	"github.com/EDDYCJY/go-example-1/routers/api"
	"github.com/EDDYCJY/go-example-1/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表：GET("/articles”)
		//获取指定文章：POST("/articles/:id”)
		//新建文章：POST("/articles”)
		//更新指定文章：PUT("/articles/:id”)
		//删除指定文章：DELETE("/articles/:id”)
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.POST("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

		//stu
		apiv1.POST("/stu/:id", v1.GetStu)
		apiv1.POST("/stus", v1.AddStu)
		apiv1.DELETE("/stu/:id", v1.DelStu)
	}
	return r
}
