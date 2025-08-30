package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LocalizationMiddleware(c *gin.Context) {
	language := "en"

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

	c.Set("lang", language)
	c.Next()
}
