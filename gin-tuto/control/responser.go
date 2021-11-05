package control

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/model"
)

type StResponser struct {
}

var Responser StResponser

func (s StResponser) RaiseResponse(c *gin.Context) error {
	var resp model.StResponse

	log.Println("in raise response")
	statusCode, exist := c.Get("responseCode")
	if exist != false {
		log.Println("exist responseCode", statusCode.(int))
		resp.Type = "OK"
		resp.Status = int64(statusCode.(int))
		resp.Detail = "alert has sent by alert server"
		resp.Title = "alert has sent"
		resp.Instance = c.FullPath()

		c.JSON(statusCode.(int), resp)

		return nil
	}
	return errors.New("code not insert gin context")
}
