package api

import (
	"github.com/gin-gonic/gin"
	"humanInfoCollection/models"
	"net/http"
)

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
	searchMap := make(map[string]string)
	searchMap["IDCard"] = id
	data := models.GetInfo(searchMap)

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"data":  data,
		"total": len(data),
	})
}

// 更新数据
func EditInfo(c *gin.Context) {

}

// 新建数据
func AddInfo(c *gin.Context) {

}
