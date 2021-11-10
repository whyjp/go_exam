package server

import (
	"webzen.com/notifyhandler/config"
	"webzen.com/notifyhandler/control/notifysender"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter(config)
	notifysender.SetConfig(config)
	r.Run(config.GetString("server.port"))
}
