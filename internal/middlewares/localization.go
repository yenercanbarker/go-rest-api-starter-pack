package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LocalizationMiddleware(c *gin.Context) {
	var language string

	languageFromQuery, isLanguageFromQueryExists := c.GetQuery("lang")
	if isLanguageFromQueryExists {
		language = languageFromQuery
		c.Next()
	}

	languageFromHeader := c.GetHeader("Accept-Language")
	if languageFromHeader != "" {
		language = languageFromHeader
		c.Next()
	}

	language = "en"
	c.Set("lang", language)
	c.Next()
}
