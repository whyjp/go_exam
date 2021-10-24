package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// Welcome godoc
// @Summary site health check is running will return "working!"
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Router /health [get]
// @Success 200
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
