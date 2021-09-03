package v1

import (
	"github.com/EDDYCJY/go-example-1/models"
	"github.com/EDDYCJY/go-example-1/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func AddStu(c *gin.Context) {
	id := c.Query("id")
	age := com.StrTo(c.Query("age")).MustInt()
	name := c.Query("name")

	data := make(map[string]interface{})
	data["id"] = id
	data["age"] = age
	data["name"] = name

	models.AddStu(data)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": make(map[string]interface{}),
	})
}

func EditStu(c *gin.Context) {
}

func DelStu(c *gin.Context) {
}

//根据条件查询
func GetStu(c *gin.Context) {

	//根据ID查询
	id := c.Param("id")
	data := models.GetStu(id)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

//获取全部
func GetStus(c *gin.Context) {

}
