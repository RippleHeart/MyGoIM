package Handlers

import (
	"github.com/gin-gonic/gin"
	"mygoim/Utils"
	"net/http"
)

func VerifyJWT(c *gin.Context) {
	var JWTRecv Utils.JWToken
	JWTRecv.Token = c.GetHeader("JWT")
	if !JWTRecv.VerifyJWT() {
		c.JSON(http.StatusOK, gin.H{"code": "101", "msg": "NO Auth!"})
		c.Abort()
		return
	}
	c.Next()
}
