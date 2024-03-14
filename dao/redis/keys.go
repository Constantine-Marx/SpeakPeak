package redis

const (
	KeyPrefix          = "speakpeak:"
	KeyPostTimeZSet    = "post:time"   // zset; 帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset; 帖子及投票时间
	KeyPostVotedZSetPF = "post:voted:" // zset; 记录用户及投票类型(1赞成0反对) KeyPostVotedZSetPrefix + post_id
	KeyCommunitySetPF  = "community:"  // set; 保存每个分区下帖子的id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
