package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	fmt.Println("Statusstatus called")
	c.Set("r_title", "working!")
	c.String(http.StatusOK, "Working!")
}
