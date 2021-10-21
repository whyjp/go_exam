// main.go
package main

import (
	"net/http"

	//"gin-tuto/docs"
	"webzen.com/notifyhandler/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @host petstore.swagger.io:8080
// @BasePath /v2
func main() {

	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/"
	docs.SwaggerInfo.BasePath = ""

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	{
		v1.GET("/health", health)
		v1.POST("/signup", signup)
		v1.POST("/login", login)
	}
	r.Use()
	r.Run()
}

// Welcome godoc
// @Summary 써머리를 직접 수정했습니다
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Router /v1/health [get]
// @Success 200
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// Welcome godoc
// @Summary 써머리를 직접 수정했습니다
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Router /v1/signup [POST]
// @Success 200
func signup(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "signed up",
	})
}

// Welcome godoc
// @Summary 써머리를 직접 수정했습니다
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Router /v1/login [POST]
// @Success 200
func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logged in",
	})
}
