package logic

import (
	"SpeakPeak/dao/mysql"
	"SpeakPeak/model"
	"SpeakPeak/pkg/snowflake"
)

func SignUp(p *model.ParamSignUp) error {
	//1.is have
	if err := mysql.CheckUserExist(p.Username); err != nil {
		//查询出错
		return err
	}
	//2.generate UID
	userID := snowflake.GenID()
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//save in mysql
	return mysql.InsertUser(user)
}

func Login(p *model.ParamLogin) (err error) {
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
