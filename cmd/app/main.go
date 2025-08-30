package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/config"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/middlewares"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/routes"
	"log"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	r.Use(middlewares.CorsMiddleware)
	r.Use(middlewares.LocalizationMiddleware)

	config.InitLogging()

	routes.InitRoutes(r)

	port := cfg.Port
	if port == "" {
		port = "1337"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
