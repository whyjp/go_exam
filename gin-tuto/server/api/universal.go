package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/control"
	"webzen.com/notifyhandler/model"
	"webzen.com/notifyhandler/util"
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
	if err := c.BindJSON(&jsonPublicMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(jsonPublicMail)

	var jsonMail model.StNotifyMail
	_, errMakeup := jsonMail.SetFrom(jsonPublicMail)
	if errMakeup != nil {
		log.Println("raise error", errMakeup)
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("errorTitle", "grafana request struct error")
		return
	}
	util.StructPrintToJson(jsonMail)
	sender := control.NewStNotifySender()
	resp, errSended := sender.SendMail(&jsonMail)
	if errSended != nil {
		log.Println(errSended)
	}
	if resp != nil {
		log.Println("resp", resp)
		c.Set("responseCode", resp.StatusCode())
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
	if err := c.BindJSON(&jsonPublicTeams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(jsonPublicTeams)

	var jsonTeams model.StNotifyTeams
	_, errMakeup := jsonTeams.SetFrom(jsonPublicTeams)
	if errMakeup != nil {
		log.Println("raise error", errMakeup)
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("errorTitle", "grafana request struct error")
		return
	}
	util.StructPrintToJson(jsonTeams)

	var sender *control.NotifySender
	resp, errSended := sender.SendTeams(&jsonTeams)
	if errSended != nil {
		log.Println(errSended)
	}
	if resp != nil {
		log.Println("resp", resp)
		c.Set("responseCode", resp.StatusCode())
	}
}
