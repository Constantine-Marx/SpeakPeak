package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("user already exists")
	ErrorUserNotExist    = errors.New("user not exists")
	ErrorInvalidPassword = errors.New("invalid password")
	ErrorUserLogged      = errors.New("user logged in elsewhere")
	ErrorInvalidID       = errors.New("invalid id")
)
