package session

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SESSION() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		if session.Get("password") == nil {
			fmt.Println("middle=============密码为空")

			c.Redirect(301,"/admin/login/index")
		}

		c.Next()
	}
}
