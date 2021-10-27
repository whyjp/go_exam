package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"webzen.com/notifyhandler/model"
)

type HealthController struct{}

// Welcome godoc
// @Summary site health check is running will return "working!"
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Router /health [get]
// @Success 200
func (h HealthController) Status(c *gin.Context) {
	log.Print("status called")
	c.String(http.StatusOK, "Working!")
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
func Param(c *gin.Context) {
	val := c.Param("test")
	action := c.Param("action")
	message := val + " " + action

	log.Println("call param", c.FullPath() == "/param/:test/*action")

	c.String(http.StatusOK, message)
}

// Welcome godoc
// @Summary jsonparam binding test
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.STNotiMail true "post jsonmail for test"
// @Router /v1/jsonMailTest [POST]
// @Success 200
func JsonMailTest(c *gin.Context) {
	var jsonMail model.STNotiMail
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

	c.JSON(resp.StatusCode(), gin.H{
		"status": "you are post put json",
		"body":   jsonMail,
		"err":    err,
	})
}
