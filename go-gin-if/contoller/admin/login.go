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
func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "admin/login/index.tmpl", gin.H{
		//"title": "Posts",
	})
}
// @Summary Get a single login
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router POST /admin/login/index
func Login(c *gin.Context) {
	appG := app.Gin{C: c}


	var m map[string]string
	m = make(map[string]string)



	session := sessions.Default(c)

	if session.Get("password") != "12346789" {
		session.Set("password", "12346789")
		session.Save()
	}

	fmt.Println(".............................", session.Get("password"))

	m["a"] = "bbb"
	m["c"] = "cccc"
	m["d"] = "dddd"
	m["e"] = "eeee"
	appG.AdminJson(http.StatusOK, e.LAYUI_SUCCESS, "成功", m)

}
