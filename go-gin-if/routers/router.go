package routers

import (
	"github.com/gin-gonic/gin"

	_ "go-gin-if/docs"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	r.Static("/layuiadmin", "./templates/layuiadmin")
	r.LoadHTMLGlob("templates/admin/**/*")

	// 后端路由
	InitAdminRouter(r)
	// 前端路由
	InitWebRouter(r)



	return r
}
