package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/control/notifysender"
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
// @Router /notify/mail [POST]
// @Success 200
func (p Universal) MailHandler(c *gin.Context) {
	var universalMail model.StUniversalProducerMail
	if err := c.BindJSON(&universalMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(universalMail)

	var jsonMail model.StNotifyMail
	_, err := jsonMail.SetFrom(universalMail)
	if err != nil {
		log.Println("raise error", err)
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("errorTitle", "grafana request struct error")
		return
	}
	util.StructPrintToJson(jsonMail)

	resp, err := notifysender.SendMail(&jsonMail)
	if err != nil {
		log.Println(err)
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
// @Router /notify/teams [POST]
// @Success 200
func (p Universal) TeamsHandler(c *gin.Context) {
	var universalTeams model.StUniversalProducerTeams
	if err := c.BindJSON(&universalTeams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(universalTeams)

	var jsonTeams model.StNotifyTeams
	_, err := jsonTeams.SetFrom(universalTeams)
	if err != nil {
		log.Println("raise error", err)
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("errorTitle", "grafana request struct error")
		return
	}
	util.StructPrintToJson(jsonTeams)

	resp, err := notifysender.SendTeams(&jsonTeams)
	if err != nil {
		log.Println(err)
	}
	if resp != nil {
		log.Println("resp", resp)
		c.Set("responseCode", resp.StatusCode())
	}
}
