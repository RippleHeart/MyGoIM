package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"mygoim/WS"
	"net/http"
)

func WSUpgrade(c *gin.Context) {
	conn, err := WS.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "WS UP ERROR!"})
		c.Abort()
		return
	}
	userID, _ := c.Get("ID")
	userName, _ := c.Get("name")
	client := &WS.Client{
		ID:   userID.(uint),
		Name: userName.(string),
		Hub:  WS.H,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	client.Hub.Register <- client

	// 每个连接两个 goroutine：读 + 写
	go client.WritePump()
	go client.ReadPump()
}
