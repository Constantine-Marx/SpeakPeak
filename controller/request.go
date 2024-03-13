package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("user not login")

const ContextUserIDKey = "userID"

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
