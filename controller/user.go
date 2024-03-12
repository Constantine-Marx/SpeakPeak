package controller

import (
	"SpeakPeak/logic"
	"SpeakPeak/model"
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
			context.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	//业务处理
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.String("username", p.Username), zap.Error(err))
		context.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
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
			context.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//2、
	if err := logic.Login(p); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg": "Wrong password",
		})
		return
	}
	//3、
	context.JSON(http.StatusOK, gin.H{
		"msg": "Login success",
	})
}
