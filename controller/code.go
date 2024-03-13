package controller

type RetCode int64

const (
	CodeSuccess RetCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMsgMap = map[RetCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "invaild param",
	CodeUserExist:       "user already exists",
	CodeUserNotExist:    "user not exists",
	CodeInvalidPassword: "invalid password",
	CodeServerBusy:      "server busy",
}

func (c RetCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if ok {
		return msg
	}
	return codeMsgMap[CodeServerBusy]
}
