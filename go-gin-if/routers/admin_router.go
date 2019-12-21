package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-gin-if/contoller/admin"
	"go-gin-if/middleware/session"
)

func InitAdminRouter(r *gin.Engine) {

	store := cookie.NewStore([]byte("secret"))




	admin_group := r.Group("/admin")
	admin_group.GET("/login/index", admin.Index)

	admin_group.Use(sessions.Sessions("admin_session", store))
	admin_group.POST("/login/login", admin.Login)

	admin_group.Use(session.SESSION())
	{
		admin_group.GET("/index/index", admin.Home)
		admin_group.GET("/index/console", admin.Console)
	}

}