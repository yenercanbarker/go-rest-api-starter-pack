package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/config"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/middlewares"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/routes"
	"io"
	"log"
	"os"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	r.Use(middlewares.CorsMiddleware)
	r.Use(middlewares.LocalizationMiddleware)

	appLogFile, _ := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(appLogFile, os.Stdout)

	infoLogFile, _ := os.OpenFile("./logs/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(io.MultiWriter(infoLogFile, os.Stdout))

	routes.InitRoutes(r)

	port := cfg.Port
	if port == "" {
		port = "1337"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
