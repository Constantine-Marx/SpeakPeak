package controller

import (
	"SpeakPeak/logic"
	"SpeakPeak/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(context *gin.Context) {
	p := new(model.Post)
	if err := context.ShouldBindJSON(p); err != nil {
		zap.L().Debug(" context.ShouldBindJSON(p)", zap.Any("err", err))
		ResponseError(context, CodeInvalidParam)
		return
	}

	userID, err := GetCurrentUserID(context)
	if err != nil {
		ResponseError(context, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//create post
	if err = logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	//resp
	ResponseSuccess(context, nil)
}

func GetPostDetailHandler(context *gin.Context) {

	//get post id
	postIDstr := context.Param("id")
	postID, err := strconv.ParseInt(postIDstr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	//get post detail
	data, err := logic.GetPostDetail(postID)
	if err != nil {
		zap.L().Error("logic.GetPostDetail failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	//resp
	ResponseSuccess(context, data)
}

func GetPostListHandler(context *gin.Context) {
	offsetStr := context.Query("offset")
	limitStr := context.Query("limit")

	var (
		limit  int64
		offset int64
		err    error
	)

	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}

	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}

	//get post list
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	//resp
	ResponseSuccess(context, data)
}
