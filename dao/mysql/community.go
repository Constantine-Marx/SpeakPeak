package mysql

import (
	"SpeakPeak/model"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func GetCommunityList() (CommunityList []*model.Community, err error) {
	sqlStr := `select community_id,community_name from community`
	CommunityList = make([]*model.Community, 0, 2)
	if err = db.Select(&CommunityList, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("GetCommunityList failed", zap.Error(err))
			err = nil
		}
	}
	return
}

func GetCommunityDetail(id int64) (community *model.CommunityDetail, err error) {
	community = new(model.CommunityDetail)
	sqlStr := `select 
			community_id, community_name, introduction, create_time
			from community 
			where community_id = ?
	`
	if err = db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
