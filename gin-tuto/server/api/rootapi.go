package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type apibase interface {
	MailHandler(c *gin.Context)
	TeamsHandler(c *gin.Context)
}
type HealthController struct{}

// Welcome godoc
// @Summary site health check is running will return "working!"
// @Description notifyhandler server heath check
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Router /health [get]
// @Success 200
func (h HealthController) Status(c *gin.Context) {
	c.Set("r_title", "working!")
	c.String(http.StatusOK, "Working!")
}
