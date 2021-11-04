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

type Grafana struct {
}

// Welcome godoc
// @Summary Grafana mail api  : have just post api
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StGrafanaAlert true "json struct for send mail"
// @Router /v1/grafana/mail [POST]
// @Success 200
func (p Grafana) MailHandler(c *gin.Context) {
	var jsonGrafana model.StGrafanaAlert
	if err := c.BindJSON(&jsonGrafana); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ := json.Marshal(jsonGrafana)
	fmt.Printf("%s \n", data)
	var jsonMail model.StNotifyMail
	fmt.Println(jsonMail)

	resp, _ := control.SendMail(&jsonMail)

	c.Set("responseCode", resp.StatusCode())
	errResp := control.Responser.RaiseResponse(c)
	if errResp != nil {
		log.Println("raise error", errResp)
	}
}

// Welcome godoc
// @Summary Grafana teams api  : have just post api
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.StGrafanaAlert true "json struct for send teams"
// @Router /v1/grafana/teams [POST]
// @Success 200
func (p Grafana) TeamsHandler(c *gin.Context) {
	var jsonGrafana model.StGrafanaAlert
	if err := c.BindJSON(&jsonGrafana); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ := json.Marshal(jsonGrafana)
	fmt.Printf("%s \n", data)
	var jsonTeams model.StNotifyTeams
	fmt.Println(jsonTeams)

	resp, _ := control.SendTeams(&jsonTeams)

	c.Set("responseCode", resp.StatusCode())
	errResp := control.Responser.RaiseResponse(c)
	if errResp != nil {
		log.Println("raise error", errResp)
	}
}
