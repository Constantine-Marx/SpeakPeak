package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code RetCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(context *gin.Context, code RetCode) {
	context.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(context *gin.Context, code RetCode, msg interface{}) {
	context.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
