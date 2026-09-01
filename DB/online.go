package DB

import (
	"context"
	"time"
)

func SetOnline(userName string) {
	var online = 1

	RDB.HSetNX(context.TODO(), "online", userName, online)
}
func SetOffline(userName string) {
	RDB.HDel(context.TODO(), "online", userName)
}
func CheckOnline(userName string) (online bool) {
	res := RDB.HGet(context.TODO(), "online", userName)
	val, err := res.Int()

	if err != nil || val == 0 {
		return false
	}
	return true
}
func SetLastOnline(userID uint) {
	MySQL.Table("users").Where("id=?", userID).Update("last_online", time.Now())
}
