package v1

import (
	"github.com/gin-gonic/gin"
	"humanInfoCollection/models"
	"humanInfoCollection/pkg/util"
	"net/http"
)

type IC struct {
	_id       string
	Name      string
	Sex       string
	Nation    string
	Birthday  string
	Hometown  string
	Career    string
	IDCard    string
	Unit      string
	Living    string
	Father    string
	Mother    string
	Brother   string
	Education string
}

// 获取所有信息
func GetInfos(c *gin.Context) {

	data := make(map[string]interface{})

	data["data"], data["total"] = models.GetInfos()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// 根据关键字获取数据
func GetInfo(c *gin.Context) {
	id := c.Param("IDCard")
	searchMap := make(map[string]interface{})
	searchMap["ID"] = id
	data := models.GetInfo(searchMap)

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"data":  data,
		"total": len(data),
	})
}

// 更新数据
func EditInfo(c *gin.Context) {
	info := &models.Info{}
	if c.BindJSON(info) == nil {

	}
	IDCard := c.Param("IDCard")
	result, err := models.EditInfo(IDCard, info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "edit fail",
			"data": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "edit success",
			"data": result,
		})
	}
}

// 新建数据
func AddInfo(c *gin.Context) {
	info := &models.Info{}
	err := c.BindJSON(info)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
	}
	result, err := models.AddInfo(info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "add success",
			"data": result,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
}

func TestInfo(c *gin.Context) {
	target := c.Param("target")
	operation := c.Query("operation")
	param := c.Query("param")

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": getMethod(target, operation)(param),
	})

}

func getMethod(target string, operation string) func(...interface{}) interface{} {
	switch target {
	case "info":
		switch operation {
		case "getJson":
			return getInfoJson
		}
	}
	return nil
}

func getInfoJson(params ...interface{}) interface{} {

	return util.GetInfo(params[0].(string), 0)
}
