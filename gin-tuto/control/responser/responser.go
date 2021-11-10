package responser

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/model"
)

type StResponser struct {
}

var Responser StResponser

func RaiseResponse(c *gin.Context) error {
	var resp model.StResponse

	log.Println("in raise response")
	statusCode, exist := c.Get("responseCode")
	if exist {
		log.Println("exist responseCode", statusCode.(int))
		resp.Type = "OK"
		resp.Status = statusCode.(int)
		resp.Detail = "alert has sent by alert server"
		errorTitle, exist := c.Get("errorTitle")
		if exist {
			resp.Title = errorTitle.(string)
		} else {
			resp.Title = "alert has sent"
		}
		resp.Instance = c.FullPath()

		c.JSON(statusCode.(int), resp)

		return nil
	}
	return errors.New("code not insert gin context")
}
