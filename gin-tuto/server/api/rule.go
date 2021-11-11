package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/model"
	"webzen.com/notifyhandler/util"
)

type Rule struct {
}

// Welcome godoc
// @Summary rule Email api  : have just post api
// @Description rule notify api for email
// @name Rule.SilentHandler
// @Accept  json
// @Produce  json
// @Param  jsonbody body model.UniversalProducerEMail true "json struct for send email"
// @Router /rule/silent [POST]
// @Success 200
func (p Rule) SilentHandler(c *gin.Context) {
	var ruleOnce model.SirentRuleOnce
	if err := c.BindJSON(&ruleOnce); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json body binding error": err.Error()})
		return
	}

	util.StructPrintToJson(ruleOnce)
	//var resultSet = processor.EMail(&ruleOnce)
	//defer util.ToContext(resultSet, c.Set)
}
