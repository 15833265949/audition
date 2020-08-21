package utils

import (
	"github.com/gin-gonic/gin"
)

type Rsp struct {
	Code int
	Msg string
	Data interface{}
}

func WriteRsp (c *gin.Context,code int,msg string,data interface{}){
	rsp := &Rsp{
		Code: code,
		Msg: msg,
		Data: data,
	}
	c.JSON(200,rsp)
}
