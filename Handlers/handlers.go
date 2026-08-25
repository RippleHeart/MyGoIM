package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type form struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var postform form
	if err := c.BindJSON(postform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

}
