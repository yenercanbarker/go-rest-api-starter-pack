package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		InitUserRoutes(api)
		InitHealthcheckRoutes(api)
		InitTestingRoutes(api)
	}
}
