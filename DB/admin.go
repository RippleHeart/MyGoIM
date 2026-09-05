package DB

import (
	"context"
	"errors"
	"time"
)

func InsertUser(name string, pwd string) error {
	var user = User{
		Name:     name,
		Password: pwd,
	}
	err := MySQL.Table("users").Create(&user).Error
	return err

}
func QueryUser(input any) (User, error) {
	var result User
	switch input.(type) {
	case uint:
		MySQL.Table("users").Select("*").Where("id=?", input).First(&result)
	case string:
		MySQL.Table("users").Select("*").Where("name=?", input).First(&result)
	default:
		return User{}, errors.New("input type error")
	}
	return result, nil
}
func QueryPwd(input string) User {
	var result User
	MySQL.Table("users").Select("password").Where("name=?", input).First(&result)
	return result
}

// UserNum 返回已注册用户总数
func UserNum() uint {
	var num uint
	MySQL.Table("users").Select("count(*)").Find(&num)
	return num
}

// OnlineNum 返回在线用户总数
func OnlineNum() uint {
	cmd := RDB.HLen(context.TODO(), "online")
	u, _ := cmd.Uint64()
	return uint(u)
}

// GetMsgNum 返回今日发帖量
func GetMsgNum() uint {
	cmd := RDB.Get(context.TODO(), "messageNum")
	num, _ := cmd.Int()
	return uint(num)
}

// GetFresherToday 返回今天新注册的用户数
func GetFresherToday() uint {
	var num uint
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 0, 1)
	MySQL.Table("user_infos").Select("count(*)").Where("created_at between ? and ?", start, end).Find(&num)
	return num
}
