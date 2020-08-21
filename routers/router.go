package routers

import (
	"github.com/gin-gonic/gin"
	"humanInfoCollection/pkg"
	v1 "humanInfoCollection/routers/api"
	"net/http"
)

func InitRouter() *gin.Engine {
	e := gin.Default()
	gin.SetMode(pkg.RunMode)

	// ping
	e.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	v1Group := e.Group("/api/v1")
	{
		v1Group.GET("/infos", v1.GetInfos)
		v1Group.GET("/infos/:IDCard", v1.GetInfo)
		v1Group.POST("/infos", v1.AddInfo)
		v1Group.PUT("/infos/:key", v1.EditInfo)
	}

	return e
}
