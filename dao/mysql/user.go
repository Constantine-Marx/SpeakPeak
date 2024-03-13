package mysql

import (
	"SpeakPeak/model"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "marxConstantine"

var (
	ErrorUserExist       = errors.New("user already exists")
	ErrorUserNotExist    = errors.New("user not exists")
	ErrorInvalidPassword = errors.New("invalid password")
)

// CheckUserExist
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser
func InsertUser(user *model.User) (err error) {
	//Encrypt
	user.Password = encryptPassword(user.Password)
	//save
	sqlStr := "insert into user(user_id, username, password) value (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(sourcePassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(sourcePassword)))
}

func Login(user *model.User) (err error) {
	srcPasssword := user.Password
	sqlStr := `select user_id,username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)

	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}

	if err != nil {
		return err
	}
	password := encryptPassword(srcPasssword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}
