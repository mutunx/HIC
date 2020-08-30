package v1

import (
	"github.com/gin-gonic/gin"
	"humanInfoCollection/models"
	"net/http"
)

func GetAdc(c *gin.Context) {
	code := c.Param("code")
	var data []models.AdministrativeDivisionsCode
	if code == "" {
		data = models.GetProvince()
	} else {
		codeType := c.Query("type")
		if codeType == "city" {
			data = models.GetCity(code)
		} else if codeType == "area" {
			data = models.GetArea(code)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
