package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"webzen.com/notifyhandler/api"
	"webzen.com/notifyhandler/docs"
	"webzen.com/notifyhandler/models"
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
// @BasePath
func NewRouter() *gin.Engine {
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080/"
	//docs.SwaggerInfo.BasePath = "v1"

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	healthroot := new(api.HealthController)
	router.GET("/health", healthroot.Status)

	v1 := router.Group("/v1")
	{
		v1.GET("/health", health)
		v1.GET("/param/:test/*action", param)
		v1.POST("/signup", signup)
		v1.POST("/login", login)
		v1.POST("/jsonTest", jsonTest)
		v1.POST("/jsonMailTest", jsonMailTest)
	}

	return router
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
// @param test path string true "test"
// @param action path string true "action"
// @Router /v1/param/{test}/{action} [get]
// @Success 200
func param(c *gin.Context) {
	val := c.Param("test")
	action := c.Param("action")
	message := val + " " + action

	fmt.Println(c.FullPath() == "/param/:test/*action")

	c.String(http.StatusOK, message)
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

// Welcome godoc
// @Summary jsonparam binding test
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body models.StJsonTest true "post json for test"
// @Router /v1/jsonTest [POST]
// @Success 200
func jsonTest(c *gin.Context) {
	var json models.StJsonTest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.First == "" || json.Second == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "json is empty"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "you are success put json",
		"info":   json,
	})
}

// Welcome godoc
// @Summary jsonparam binding test
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body models.StMailTest true "post jsonmail for test"
// @Router /v1/jsonMailTest [POST]
// @Success 200
func jsonMailTest(c *gin.Context) {
	var jsonMail_ models.StMailTest
	if err := c.ShouldBindJSON(&jsonMail_); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//if jsonMail.from == "" || jsonMail.to == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{"status": "json is empty"})
	//	return
	//}

	fmt.Println(jsonMail_.From)
	pbytes, _ := json.Marshal(jsonMail_)
	buff := bytes.NewBuffer(pbytes)

	// Request 객체 생성
	req, err := http.NewRequest("POST", "http://10.105.33.38/alert/api/v2/email", buff)
	if err != nil {
		panic(err)
	}

	//Content-Type 헤더 추가
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
		c.String(http.StatusOK, str)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "you are success put json",
		"info":   jsonMail_,
	})
}
