package controller

import (
	"SpeakPeak/dao/mysql"
	"SpeakPeak/logic"
	"SpeakPeak/model"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(context *gin.Context) {
	//参数校验
	var p model.ParamSignUp
	if err := context.ShouldBindJSON(&p); err != nil {
		//error
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断是不是validator类型错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(context, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(context, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(context, CodeUserExist)
			return
		}
		ResponseError(context, CodeServerBusy)
		return
	}
	//返回响应
	context.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})

}

func LoginHandler(context *gin.Context) {
	//参数校验
	p := new(model.ParamLogin)
	if err := context.ShouldBindJSON(&p); err != nil {
		//error
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断是不是validator类型错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(context, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(context, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2、
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(context, CodeUserNotExist)
			return
		}
		ResponseError(context, CodeInvalidPassword)
	}
	//3、
	ResponseSuccess(context, nil)
}
