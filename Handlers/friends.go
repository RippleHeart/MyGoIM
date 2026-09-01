package Handlers

import (
	"github.com/gin-gonic/gin"
	"mygoim/DB"
	"net/http"
)

func GetFriends(c *gin.Context) {
	userID, _ := c.Get("ID")
	result := DB.QueryFrdAll(userID.(uint))
	c.JSON(http.StatusOK, gin.H{"code": "003", "msg": "successfully!", "friends": result})
}

func AddFriend(c *gin.Context) {
	userID, _ := c.Get("ID")
	frdName := c.Param("friendname")
	frd, err := DB.QueryUser(frdName)
	//对方是否存在，是否为自己
	if frd.ID == 0 || userID == frd.ID || err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Illegal Addition!"})
		return
	}
	//是否已添加
	result := DB.QueryFrd(userID.(uint), frd.ID)
	if result.Name != "" {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Repeated Addition!"})
		return
	}
	//插入到中间表
	err = DB.InsertFrd(userID.(uint), frd.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "002", "msg": "Successful Addition!"})
}
