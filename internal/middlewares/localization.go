package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LocalizationMiddleware(c *gin.Context) {
	language := "en"

	languageFromQuery, isLanguageFromQueryExists := c.GetQuery("lang")
	if isLanguageFromQueryExists {
		language = languageFromQuery
	}

	languageFromHeader := c.GetHeader("Locale")
	if languageFromHeader != "" {
		language = languageFromHeader
	}

	c.Set("lang", language)
	c.Next()
}
