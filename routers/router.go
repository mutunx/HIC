package routers

import (
	"github.com/gin-gonic/gin"
	"humanInfoCollection/pkg"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	e := gin.Default()
	gin.SetMode(pkg.RunMode)

	v1Group := e.Group("/api/v1")
	{
		v1Group.GET("/test", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"time": time.Now(),
				"msg":  "test success ! Hello you",
			})
		})
	}

	return e
}
