package server

import (
	"webzen.com/notifyhandler/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter(config)
	r.Run(config.GetString("server.port"))
}
