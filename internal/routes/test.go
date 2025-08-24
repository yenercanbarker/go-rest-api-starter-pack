package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitTestingRoutes(router *gin.RouterGroup) {
	router.GET("/testing", func(c *gin.Context) {
		testingValues := []string{"first", "second", "third"}

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"data":   testingValues,
		})
	})
}
