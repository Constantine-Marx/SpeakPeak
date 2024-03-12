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
	//手动对请求参数进行业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 {
	//	//error
	//	zap.L().Error("SignUp with invalid param")
	//	context.JSON(http.StatusOK, gin.H{
	//		"msg": "error in request params",
	//	})
	//	return
	//}

	//业务处理
	logic.SignUp(&p)
	//返回响应
	context.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})

}
