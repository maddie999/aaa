package app

import (
	"github.com/gin-gonic/gin"

	"go-gin-if/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

type ResponseaAdminjson struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) AdminJson(httpCode, errCode int, msg interface{}, data interface{}) {
	g.C.JSON(httpCode, ResponseaAdminjson{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}
