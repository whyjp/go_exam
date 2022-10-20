package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	log.Println("config init")
	var err error
	config = viper.New()
	config.SetConfigType("toml")
	config.SetConfigName(env)
	config.AddConfigPath("config/")
	config.AddConfigPath("./")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing config file")
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
