package logic

import (
	"SpeakPeak/dao/mysql"
	"SpeakPeak/model"
)

func GetCommunityList() (CommunityList []*model.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (CommunityDetail *model.CommunityDetail, err error) {
	return mysql.GetCommunityDetail(id)
}
