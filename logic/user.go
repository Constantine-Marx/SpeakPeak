package logic

import (
	"SpeakPeak/dao/mysql"
	"SpeakPeak/model"
)

func SignUp(p *model.ParamSignUp) {
	//is have
	mysql.QueryUserByUsername()
	//generate UID

	//save in mysql
}
