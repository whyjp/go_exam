package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/control/processor"
	"webzen.com/notifyhandler/model"
	"webzen.com/notifyhandler/util"
)

type Universal struct {
}

// Welcome godoc
// @Summary universal Email api  : have just post api
// @Description universal notify api for email
// @name Universal.EMailHandler
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StUniversalProducerEMail true "json struct for send email"
// @Router /notify/email [POST]
// @Success 200
func (p Universal) EMailHandler(c *gin.Context) {
	var universalEMail model.StUniversalProducerEMail
	if err := c.BindJSON(&universalEMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(universalEMail)
	var resultSet = processor.EMail(&universalEMail)
	defer util.ToContext(resultSet, c.Set)
}

// Welcome godoc
// @Summary Grafana teams api  : have just post api
// @Description universal notify api for teams
// @name Universal.TeamsHandler
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

	var resultSet = processor.Teams(&universalTeams)
	defer util.ToContext(resultSet, c.Set)
}
