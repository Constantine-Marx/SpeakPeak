package logic

import (
	"SpeakPeak/dao/mysql"
	"SpeakPeak/dao/redis"
	"SpeakPeak/model"
	"SpeakPeak/pkg/snowflake"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func CreatePost(p *model.Post) (err error) {
	//gen id
	p.ID = int64(snowflake.GenID())
	//save db
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	if err != nil {
		return err
	}
	return err
}

func GetPostDetail(postID int64) (post model.ApiPostDetail, err error) {
	//get post detail
	postData, err := mysql.GetPostByID(postID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID failed",
			zap.Int64("pid", postID),
			zap.Error(err))
		return post, err
	}
	//get author name
	authName, errs := mysql.GetUsernameByUserID(postData.AuthorID)
	if errs != nil {
		return post, errs
	}
	post.AuthorName = authName
	//get community name
	community, err := mysql.GetCommunityBYID(postData.CommunityID)
	if err != nil {
		return post, err
	}

	post.Post = postData
	post.Community = community
	return
}

func GetPostList(offset, limit int64) (data []*model.ApiPostDetail, err error) {
	//get post list
	postList, err := mysql.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("mysql.GetPostList failed", zap.Error(err))
		return nil, err
	}
	data = make([]*model.ApiPostDetail, 0, len(postList))
	//get author name
	for _, v := range postList {
		authName, errs := mysql.GetUsernameByUserID(v.AuthorID)
		if errs != nil {
			if errors.Is(errs, sql.ErrNoRows) {
				zap.L().Warn("mysql.GetUsernameByUserID failed", zap.Error(errs))
				continue
			} else {
				return nil, errs
			}
		}
		//get community name
		community, err := mysql.GetCommunityBYID(v.CommunityID)
		if err != nil {
			return nil, err
		}
		post := model.ApiPostDetail{
			AuthorName: authName,
			Post:       v,
			Community:  community,
		}
		data = append(data, &post)
	}
	return
}
