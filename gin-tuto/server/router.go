package server

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"webzen.com/notifyhandler/control/responser"
	"webzen.com/notifyhandler/docs"
	"webzen.com/notifyhandler/server/api"
)

// @title Webzen NotifyHandler server
// @version 1.0
// @description service based notifyhandler server it use external endpoint notify server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email youngjoopark@webzen.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080/
// @BasePath
func NewRouter(config *viper.Viper) *gin.Engine {
	logName := config.GetString("server.log")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fileLog, err := os.OpenFile(exPath+"/"+logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(fileLog, os.Stdout)

	runtime.GOMAXPROCS(runtime.NumCPU())

	docs.SwaggerInfo.Title = "NotifyHandler API"
	docs.SwaggerInfo.Description = "This service is used to notify and notify managing."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/"

	router := gin.New()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	healthroot := new(api.HealthController)
	router.GET("/health", healthroot.Status)
	notify := router.Group("/notify")
	notify.Use(responseMiddleware())
	{
		pub := new(api.Universal)
		notify.POST("/email", pub.EMailHandler)
		notify.POST("/teams", pub.TeamsHandler)

		grafanaRouter := notify.Group("grafana")
		{
			grafana := new(api.Grafana)
			grafanaRouter.POST("/email", grafana.EMailHandler)
			grafanaRouter.POST("/teams", grafana.TeamsHandler)
		}
	}
	return router
}

func responseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		errResp := responser.RaiseResponse(ctx)
		if errResp != nil {
			log.Println("raise error", errResp)
		}
	}
}
