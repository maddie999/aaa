package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-if/pkg/app"
	"go-gin-if/pkg/e"
	"net/http"
)

// @Summary Get a single index
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router GET /api/v1/index
func Index(c *gin.Context) {
	appG := app.Gin{C: c}


	appG.Response(http.StatusOK, e.SUCCESS, 11)
}
