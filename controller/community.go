package controller

import (
	"SpeakPeak/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(context *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}

func CommunityDetailHandler(context *gin.Context) {
	//1.get community id
	target_id := context.Param("id")
	id, err := strconv.ParseInt(target_id, 10, 64)
	if err != nil {
		ResponseError(context, CodeInvalidParam)
		return
	}
	//2.get data
	data, err := logic.GetCommunityDetail(id)
	//3.return
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}
