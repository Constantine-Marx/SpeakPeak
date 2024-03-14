package model

//define request struct

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token"`
}

type ParamVoteData struct {
	PostID int64 `json:"post_id,string" binding:"required"`             // 帖子id
	Direct int8  `json:"direct,string" binding:"required,oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)
}
