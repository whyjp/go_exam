package server

import (
	"io"
	"log"
	"os"
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"webzen.com/notifyhandler/api"
	"webzen.com/notifyhandler/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080/
// @BasePath
func NewRouter(config *viper.Viper) *gin.Engine {
	logName := config.GetString("server.log")
	fileLog, err := os.Create(logName)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(fileLog, os.Stdout)

	runtime.GOMAXPROCS(runtime.NumCPU())

	docs.SwaggerInfo.Title = "NotificationHandler API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/"
	//docs.SwaggerInfo.BasePath = "v1"

	router := gin.New()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())

	//router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		healthroot := new(api.HealthController)
		router.GET("/health", healthroot.Status)
		v1.GET("/param/:test/*action", api.Param)
		v1.POST("/jsonMailTest", api.JsonMailTest)
	}
	return router
}
