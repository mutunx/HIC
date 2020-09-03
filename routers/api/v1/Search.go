package v1

import (
	"github.com/gin-gonic/gin"
	"humanInfoCollection/models"
	"net/http"
)

func Search(c *gin.Context) {
	target := c.Param("target")
	var data interface{}
	switch target {
	case "info":
		data = searchInfo(c)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func searchInfo(c *gin.Context) interface{} {
	keyword := c.Query("keyword")
	if keyword != "" {
		return models.SearchInfo(keyword)
	} else if c.Query("sTime") != "" {
		sTime := c.Query("sTime")
		eTime := c.Query("eTime")
		return models.SearchInfoBySpan(sTime, eTime)
	}
	return nil
}
