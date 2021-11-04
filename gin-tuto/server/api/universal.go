package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/control"
	"webzen.com/notifyhandler/model"
)

type Universal struct {
}

// Welcome godoc
// @Summary universal mail api  : have just post api
// @Description universal notify api for mail
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StUniversalProducerMail true "json struct for send mail"
// @Router /v1/mail [POST]
// @Success 200
func (p Universal) MailHandler(c *gin.Context) {
	var jsonPublicMail model.StUniversalProducerMail
	if err := c.ShouldBindJSON(&jsonPublicMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ := json.Marshal(jsonPublicMail)
	fmt.Printf("%s \n", data)
	var jsonMail model.StNotifyMail
	fmt.Println(jsonMail)

	resp, errSended := control.SendMail(&jsonMail)
	if errSended != nil {
		log.Println(errSended)
	}
	if resp != nil {
		log.Println("resp", resp)
		c.Set("responseCode", resp.StatusCode())
		errResp := control.Responser.RaiseResponse(c)
		if errResp != nil {
			log.Println("raise error", errResp)
		}
	}
}

// Welcome godoc
// @Summary Grafana teams api  : have just post api
// @Description universal notify api for teams
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StUniversalProducerTeams true "json struct for send teams"
// @Router /v1/teams [POST]
// @Success 200
func (p Universal) TeamsHandler(c *gin.Context) {
	var jsonPublicTeams model.StUniversalProducerTeams
	if err := c.ShouldBindJSON(&jsonPublicTeams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ := json.Marshal(jsonPublicTeams)
	fmt.Printf("%s \n", data)
	var jsonTeams model.StNotifyTeams
	fmt.Println(jsonTeams)

	resp, errSended := control.SendTeams(&jsonTeams)
	if errSended != nil {
		log.Println(errSended)
	}
	if resp != nil {
		log.Println("resp", resp)
		c.Set("responseCode", resp.StatusCode())
		errResp := control.Responser.RaiseResponse(c)
		if errResp != nil {
			log.Println("raise error", errResp)
		}
	}
}
