package server

import (
	"webzen.com/notifyhandler/config"
	"webzen.com/notifyhandler/control/notifysender"
)

func Init() {
	config := config.GetConfig()
	notifysender.SetConfig(config)
	r := NewRouter(config)
	r.Run(config.GetString("server.port"))
}
