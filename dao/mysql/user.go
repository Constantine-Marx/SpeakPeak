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
	ErrorUserLogged      = errors.New("user logged in elsewhere")
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

func CheckUserToken(user *model.User, token string) (err error) {
	sqlStr := `select count(user_id) from user_token where user_id = ?`
	var count int
	if err = db.Get(&count, sqlStr, user.UserID); err != nil {
		return err
	}
	if count == 0 {
		//Insert
		sqlStr = "insert into user_token(user_id, token) value (?,?)"
		_, err = db.Exec(sqlStr, user.UserID, token)
	} else if count > 0 {
		//Check if the token is the same
		sqlStr = `select token from user_token where user_id = ?`
		var GetToken string
		err = db.Get(&GetToken, sqlStr, user.UserID)
		if err != nil {
			return err
		}
		if GetToken == token {
			//Update
			sqlStr = "update user_token set token = ? where user_id = ?"
			_, err = db.Exec(sqlStr, token, user.UserID)
			return nil
		} else {
			return ErrorUserLogged
		}
	}
	return nil
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
