package Service

import (
	"github.com/gin-gonic/gin"
	"mygoim/DB"
	"mygoim/Utils"
	"net/http"
)

func VerifyJWT(c *gin.Context) {
	var JWTRecv Utils.JWToken
	JWTRecv.Token = c.GetHeader("JWT")
	username, ok := JWTRecv.VerifyJWT()
	nameParam := c.Param("name")
	if !ok || nameParam != username {
		c.JSON(http.StatusOK, gin.H{"code": "101", "msg": "NO Auth!"})
		c.Abort()
		return
	}
	user, err := DB.QueryUser(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "101", "msg": err.Error()})
		c.Abort()
		return
	}
	c.Set("ID", user.ID)
	c.Set("name", username)
	c.Next()
}
