// main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"webzen.com/notifyhandler/config"
	"webzen.com/notifyhandler/server"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Printf("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	server.Init()
}
