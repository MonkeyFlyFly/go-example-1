package routers

import (
	"github.com/EDDYCJY/go-example-1/pkg/setting"
	v1 "github.com/EDDYCJY/go-example-1/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
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
	}
	return r
}
