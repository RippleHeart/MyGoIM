package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygoim/DB"
	"net/http"
)

func GetFriends(c *gin.Context) {
	userID, _ := c.Get("ID")
	var results []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	DB.MySQL.
		Raw("select b.name, b.id from users a,users b,user_users c where a.id=? and ((a.id=c.active_id and b.id =c.passive_id) or  (b.id=c.active_id and a.id =c.passive_id))", userID).
		Find(&results)
	fmt.Println(results)
	c.JSON(http.StatusOK, gin.H{"code": "003", "msg": "successfully!", "friends": results})
}

func AddFriend(c *gin.Context) {
	activeID, _ := c.Get("ID")
	passivename := c.Query("name")
	var TempUser struct {
		ID uint
	}
	var count int64
	DB.MySQL.Table("users").Select("id").Where("name=?", passivename).First(&TempUser)
	if TempUser.ID == 0 || activeID == TempUser.ID {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Illegal Addition!"})
		return
	}
	DB.MySQL.Table("user_users").
		Where("((active_id = ? AND passive_id = ?) OR (active_id = ? AND passive_id = ?)) AND Break = ?",
			activeID, TempUser.ID, TempUser.ID, activeID, "false").
		Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Repeated Addition!"})
		return
	}
	fmt.Printf("aID=%d,pID=%d", activeID, TempUser.ID)
	// 插入好友申请
	err := DB.MySQL.Table("user_users").Create(&DB.UserUser{
		ActiveID:  activeID.(uint),
		PassiveID: TempUser.ID,
		Break:     false,
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "002", "msg": "Successful Addition!"})
}
