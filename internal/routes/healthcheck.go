package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitHealthcheckRoutes(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Server is ready",
		})
	})
}
