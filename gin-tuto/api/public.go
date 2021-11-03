package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/control"
	"webzen.com/notifyhandler/model"
)

type Public struct {
}

// Welcome godoc
// @Summary public mail api  : have just post api
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StPublicProducerMail true "json struct for send mail"
// @Router /v1/public/mail [POST]
// @Success 200
func (p Public) MailHandler(c *gin.Context) {
	var jsonPublicMail model.StPublicProducerMail
	if err := c.ShouldBindJSON(&jsonPublicMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ := json.Marshal(jsonPublicMail)
	fmt.Printf("%s \n", data)
	var jsonMail model.StNotifyMail
	fmt.Println(jsonMail)

	resp, _ := control.SendMail(&jsonMail)

	result, _ := control.Responser.GetResult(resp.StatusCode(), c) // model.StResponse

	c.JSON(resp.StatusCode(), result)
}

// Welcome godoc
// @Summary Grafana teams api  : have just post api
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StPublicProducerTeams true "json struct for send teams"
// @Router /v1/public/teams [POST]
// @Success 200
func (p Public) TeamsHandler(c *gin.Context) {
	var jsonPublicTeams model.StPublicProducerTeams
	if err := c.ShouldBindJSON(&jsonPublicTeams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ := json.Marshal(jsonPublicTeams)
	fmt.Printf("%s \n", data)
	var jsonTeams model.StNotifyTeams
	fmt.Println(jsonTeams)

	resp, _ := control.SendTeams(&jsonTeams)

	result, _ := control.Responser.GetResult(resp.StatusCode(), c) // model.StResponse

	c.JSON(resp.StatusCode(), result)
}
