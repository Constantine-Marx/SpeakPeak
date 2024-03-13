package controller

type RetCode int64

const (
	CodeSuccess RetCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeUserLogged
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
	CodeInvalidAuth
	CodeInvalidTokenFormat
	CodeTokenExpired
)

var codeMsgMap = map[RetCode]string{
	CodeSuccess:            "success",
	CodeInvalidParam:       "invaild param",
	CodeUserExist:          "user already exists",
	CodeUserNotExist:       "user not exists",
	CodeUserLogged:         "user logged in elsewhere",
	CodeInvalidPassword:    "invalid password",
	CodeServerBusy:         "server busy",
	CodeInvalidToken:       "invalid token",
	CodeNeedLogin:          "need login",
	CodeInvalidAuth:        "invalid auth",
	CodeInvalidTokenFormat: "invalid token format",
	CodeTokenExpired:       "token expired",
}

func (c RetCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if ok {
		return msg
	}
	return codeMsgMap[CodeServerBusy]
}
