package logic

import (
	"SpeakPeak/dao/redis"
	"SpeakPeak/model"
	"strconv"

	"go.uber.org/zap"
)

func VoteForPost(userID int64, p *model.ParamVoteData) (err error) {
	zap.L().Debug("VoteForPost", zap.Any("userID", userID), zap.Any("p", p), zap.Any("p.PostID", p.PostID), zap.Any("p.Direct", p.Direct))
	return redis.VoteForPost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Direct))
}
