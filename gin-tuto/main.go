// main.go
package main

import (
	"net/http"

	//"gin-tuto/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerfiles"
)

func main() {
	/*
		docs.SwaggerInfo.Title = "Swagger API"
		docs.SwaggerInfo.Description = "This is a sample server for Swagger."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "petstore.swagger.io"
		docs.SwaggerInfo.BasePath = "/v2"
	*/
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
// @Param name path string true "User name"
// @Router /v1/{name} [get]
// @Success 200
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func signup(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "signed up",
	})
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logged in",
	})
}
