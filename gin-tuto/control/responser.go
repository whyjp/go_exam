package control

import (
	"github.com/gin-gonic/gin"
	"webzen.com/notifyhandler/model"
)

type StResponser struct {
}

var Responser StResponser

func (s StResponser) MakeResponse(statusCode int, c *gin.Context) (*model.StResponse, error) {
	var resp *model.StResponse

	resp.Type = "OK"
	resp.Status = int64(statusCode)
	resp.Detail = "alert has sent by alert server"
	resp.Title = "alert has sent"
	resp.Instance = c.FullPath()

	return resp, nil
}
