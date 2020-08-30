package routers

import (
	"github.com/gin-gonic/gin"
	"humanInfoCollection/pkg"
	"humanInfoCollection/routers/api"
	"humanInfoCollection/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	e := gin.Default()

	gin.SetMode(pkg.RunMode)

	// ping
	e.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	e.GET("/auth", api.GetAuth)

	v1Group := e.Group("/api/v1")
	//v1Group.Use(jwt.JWT())
	{
		v1Group.GET("/infos", v1.GetInfos)
		v1Group.GET("/infos/:IDCard", v1.GetInfo)
		v1Group.POST("/infos", v1.AddInfo)
		v1Group.PUT("/infos/:IDCard", v1.EditInfo)
		v1Group.GET("/adc/:code", v1.GetAdc)
		v1Group.GET("/adc", v1.GetAdc)
		v1Group.GET("/test/:target", v1.TestInfo)
	}

	return e
}
