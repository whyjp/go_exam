package server

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	logName := config.GetString("server.log")
	fileLog, _ := os.Create(logName)
	gin.DefaultWriter = io.MultiWriter(fileLog)
	r.Run(config.GetString("server.port"))
}
