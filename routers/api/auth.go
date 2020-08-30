package api

import (
	"github.com/gin-gonic/gin"
	"humanInfoCollection/models"
	"humanInfoCollection/pkg/e"
	"humanInfoCollection/pkg/util"
	"net/http"
)

func GetAuth(c *gin.Context) {
	loginName := c.Query("loginName")
	password := c.Query("password")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	// 验证密码
	isExist := models.CheckAuth(loginName, password)
	if isExist {
		// 获取token
		token, er := util.GenerateToken(loginName, password)
		if er != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = e.SUCCESS
		}

	} else {
		code = e.ERROR_AUTH
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
