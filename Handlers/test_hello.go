package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": "999", "msg": "Hello!"})
}
