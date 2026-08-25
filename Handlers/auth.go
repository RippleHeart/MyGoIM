package Handlers

import (
	"github.com/gin-gonic/gin"
	"mygoim/DB"
	"net/http"
)

type form struct {
	Username string `gorm:"column:name" json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var postform form
	var virify form
	if err := c.BindJSON(&postform); err != nil || postform.Username == "" || postform.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  "101",
			"error": "请输入",
		})
		return
	}
	//todo 加解密密码比较
	DB.MySQL.Table("users").Select("name", "password").Where("name=?", postform.Username).First(&virify)
	if virify.Password != postform.Password || virify.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "101",
			"msg":  "登录失败",
		})
		return
	}
	//todo  返回token
	c.JSON(http.StatusOK, gin.H{
		"code":  "100",
		"msg":   "登录成功",
		"token": "",
	})

}

func Register(c *gin.Context) {
	var postform form
	if err := c.BindJSON(&postform); err != nil || postform.Username == "" || postform.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  "101",
			"error": "请输入",
		})
		return
	}
	//todo  加密密码
	err := DB.MySQL.Table("users").Create(&postform).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  "101",
			"error": "数据库出错",
		})
		return
	}
	//todo 返回token
	c.JSON(http.StatusOK, gin.H{
		"code":  "100",
		"msg":   "登录成功",
		"token": "",
	})

}
