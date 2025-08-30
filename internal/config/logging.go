package config

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func InitLogging() {
	appLogFile, _ := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(appLogFile, os.Stdout)

	infoLogFile, _ := os.OpenFile("./logs/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(io.MultiWriter(infoLogFile, os.Stdout))
}
