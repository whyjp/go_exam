package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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
	v1.Use(gin.Logger())
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
// @Param  jsonbody body models.STNotiMail true "post jsonmail for test"
// @Router /v1/jsonMailTest [POST]
// @Success 200
func jsonMailTest(c *gin.Context) {
	var jsonMail models.STNotiMail
	if err := c.ShouldBindJSON(&jsonMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	type AuthSuccess struct {
	}
	type AuthError struct {
	}

	fmt.Println(jsonMail)

	client := resty.New()
	resp, err := client.R().
		SetBody(jsonMail).
		SetResult(AuthSuccess{}). // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}).   // or SetError(AuthError{}).
		Post("http://10.105.33.38/alert/api/v2/email")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	//fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	c.JSON(http.StatusOK, gin.H{
		"status": "you are post put json",
		"info":   jsonMail,
		"err":    err,
	})
}
