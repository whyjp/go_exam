package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/control/center"
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
	var resultSet = center.ProcessMail(&universalMail)
	defer util.ToContext(resultSet, c.Set)
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

	var resultSet = center.ProcessTeams(&universalTeams)
	defer util.ToContext(resultSet, c.Set)
}
