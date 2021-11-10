package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/control/center"
	"webzen.com/notifyhandler/model"
	"webzen.com/notifyhandler/util"
)

type Grafana struct {
}

// Welcome godoc
// @Summary Grafana mail api  : have just post api
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StGrafanaAlert true "json struct for send mail"
// @Router /notify/grafana/mail [POST]
// @Success 200
func (p Grafana) MailHandler(c *gin.Context) {
	var jsonGrafana model.StGrafanaAlert
	if err := c.BindJSON(&jsonGrafana); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(jsonGrafana)
	resultSet := center.ProcessMail(&jsonGrafana)
	defer util.ToContext(resultSet, c.Set)
}

// Welcome godoc
// @Summary Grafana teams api  : have just post api
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StGrafanaAlert true "json struct for send teams"
// @Router /notify/grafana/teams [POST]
// @Success 200
func (p Grafana) TeamsHandler(c *gin.Context) {
	var jsonGrafana model.StGrafanaAlert
	if err := c.BindJSON(&jsonGrafana); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(jsonGrafana)
	resultSet := center.ProcessTeams(&jsonGrafana)
	defer util.ToContext(resultSet, c.Set)
}
