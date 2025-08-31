package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/utils"
	"net/http"
)

func InitTestingRoutes(router *gin.RouterGroup) {
	router.GET("/testing", func(c *gin.Context) {
		testingValues := []string{"first", "second", "third"}

		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": utils.Translate(c, "hello", nil),
			"messageWithValues": utils.Translate(c, "welcome", map[string]string{
				"Name": "Yenercan",
			}),
			"messageInDepth": utils.Translate(c, "errors.internalServer", nil),
			"data":           testingValues,
		})
	})
}
