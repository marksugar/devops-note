package router

import (
	"etcd/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == "Release" {
		gin.SetMode(gin.ReleaseMode)
	}
	// r := gin.New()
	r := gin.Default()

	r.POST("/tables", controllers.Api.Tables)
	r.POST("/adds", controllers.Api.Adds)
	r.PUT("/updates/:id", controllers.Api.Updates)
	r.DELETE("/dels/:id", controllers.Api.Dels)
	r.PUT("/upid/:id", controllers.Api.Updates2)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
