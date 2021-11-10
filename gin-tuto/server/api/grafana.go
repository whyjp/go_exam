package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/control/processor"
	"webzen.com/notifyhandler/model"
	"webzen.com/notifyhandler/util"
)

type Grafana struct {
}

// Welcome godoc
// @Summary Grafana mail api  : have just post api
// @Description 그라파나로 부터 메일을 통해 메세지를 보내고자할 때 사용 합니다
// @name Grafana.EMailHandler
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StGrafanaAlert true "json struct for send email"
// @Router /notify/grafana/email [POST]
// @Success 200
func (p Grafana) EMailHandler(c *gin.Context) {
	var jsonGrafana model.StGrafanaAlert
	if err := c.BindJSON(&jsonGrafana); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(jsonGrafana)
	resultSet := processor.EMail(&jsonGrafana)
	defer util.ToContext(resultSet, c.Set)
}

// Welcome godoc
// @Summary Grafana teams api  : have just post api
// @Description 그라파나로 부터 팀즈 를 향해 메세지를 보내고자 할 때 사용 합니다
// @name Grafana.TeamsHandler
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
	resultSet := processor.Teams(&jsonGrafana)
	defer util.ToContext(resultSet, c.Set)
}
