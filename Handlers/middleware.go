package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygoim/DB"
	"mygoim/Utils"
	"net/http"
)

func VerifyJWT(c *gin.Context) {
	var JWTRecv Utils.JWToken
	var userID struct {
		ID uint
	}
	JWTRecv.Token = c.GetHeader("JWT")
	username, ok := JWTRecv.VerifyJWT()
	nameParam := c.Param("name")
	if !ok || nameParam != username {
		c.JSON(http.StatusOK, gin.H{"code": "101", "msg": "NO Auth!"})
		c.Abort()
		return
	}

	DB.MySQL.Table("users").Select("id").Where("name = ?", username).First(&userID)
	fmt.Println(userID)
	c.Set("ID", userID.ID)
	c.Set("name", username)
	c.Next()
}
