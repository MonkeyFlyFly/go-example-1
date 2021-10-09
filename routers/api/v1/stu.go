package v1

import (
	"encoding/json"
	"fmt"
	"github.com/EDDYCJY/go-example-1/models"
	"github.com/EDDYCJY/go-example-1/pkg/e"
	"github.com/EDDYCJY/go-example-1/pkg/gredis"
	"github.com/EDDYCJY/go-example-1/pkg/logging"
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
	id := com.StrTo(c.Param("id")).MustInt()
	models.DeleteStu(id)
}

//根据条件查询
func GetStu(c *gin.Context) {
	//appG := app.Gin{c}
	//根据ID查询
	id := c.Param("id")
	//valid := validation.Validation{}
	//valid.Min(id, 1, "id").Message("ID必须大于0")

	//if valid.HasErrors() {
	//	app.MarkErrors(valid.Errors)
	//	appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	//	return
	//}

	exists := gredis.Exists(com.ToStr(id))

	fmt.Println(exists, "----------exists--------")
	var data interface{}
	if exists {
		datas, _ := gredis.Get(id)
		var cacheStu *models.Stu
		json.Unmarshal(datas, &cacheStu) //redis返回对象处理
		fmt.Println(cacheStu)
		fmt.Println(data, "----------redis--------")
	} else {
		data = models.GetStu(id)
		_ = gredis.Set(id, data, 10000)
		fmt.Println(data, "------------db------")
	}

	logging.Info(data, "------------------")

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

//获取全部
func GetStus(c *gin.Context) {

}
