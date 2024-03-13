package logic

import (
	"SpeakPeak/dao/mysql"
	"SpeakPeak/model"
	"SpeakPeak/pkg/jwt"
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

func Login(p *model.ParamLogin) (token string, err error) {
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
		Token:    p.Token,
	}

	if err = mysql.Login(user); err != nil {
		return "", err
	}
	token, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return "", err
	}
	err = mysql.CheckUserToken(user, token)
	if err != nil {
		return "", err
	}
	return
}
