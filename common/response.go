package common

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func ReturnResponse(ctx *gin.Context, code int, msg string, body interface{}) {
	res := &Response{Code: code, Message: msg, Body: body}
	ctx.JSON(200, res)
}
