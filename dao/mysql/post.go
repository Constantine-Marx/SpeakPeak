package mysql

import (
	"SpeakPeak/model"

	"go.uber.org/zap"
)

func CreatePost(p *model.Post) (err error) {
	sqlStr := `insert into post(
    post_id, title, content, author_id, community_id)
    values (?,?,?,?,?)`

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostByID(id int64) (post *model.Post, err error) {
	post = new(model.Post)
	sqlStr := `select 
		post_id, title, content, author_id, community_id, status, create_time
		from post
		where post_id = ?`
	err = db.Get(post, sqlStr, id)
	return
}

func GetUsernameByUserID(id int64) (username string, err error) {
	sqlStr := `select username from user where user_id = ?`
	err = db.Get(&username, sqlStr, id)
	return
}

func GetCommunityBYID(id int64) (community *model.Community, err error) {
	community = new(model.Community)
	sqlStr := `select community_id, community_name from community where community_id = ?`
	err = db.Get(community, sqlStr, id)
	return
}

func GetPostList(limit, offset int64) (data []*model.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, status, create_time
		from post
		order by create_time desc
		limit ? offset ?`
	data = make([]*model.Post, 0, 2)
	err = db.Select(&data, sqlStr, limit, offset)
	if err != nil {
		zap.L().Error("db.Select(data, sqlStr) failed", zap.Error(err))
		return nil, err
	}
	return
}
