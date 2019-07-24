package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-scaffold/pkgs/e"
)

type Gin struct {
	C *gin.Context
}

// 返回数据结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 填充返回报文信息
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	fmt.Println(e.GetMsg(errCode))
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
