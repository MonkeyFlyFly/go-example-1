package v1

import (
	"github.com/EDDYCJY/go-example-1/models"
	"github.com/EDDYCJY/go-example-1/pkg/e"
	"github.com/EDDYCJY/go-example-1/pkg/setting"
	"github.com/EDDYCJY/go-example-1/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} string "ok"
// @Failure 500 {object} string "not ok"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)

	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
