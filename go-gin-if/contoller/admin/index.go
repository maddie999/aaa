package admin

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gin-if/pkg/app"
	"go-gin-if/pkg/e"
	"net/http"
)

// @Summary Get a single index
// @Produce  html
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router GET /admin/login/index
func Home(c *gin.Context) {

	session := sessions.Default(c)

	fmt.Println("-----------------------------", session.Get("password"))

	c.HTML(http.StatusOK, "admin/index/index.tmpl", gin.H{
		"title": "Posts",
	})
}

func Console(c *gin.Context) {

	session := sessions.Default(c)

	appG := app.Gin{C: c}
	appG.AdminJson(http.StatusOK, e.LAYUI_SUCCESS, "成功", session.Get("password"))
}