package Handlers

import (
	"fmt"
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
	ch, err := WS.NewChannel()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
		c.Abort()
		return
	}
	q, err := ch.QueueDeclare(userName.(string), true, false, false, false, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
		c.Abort()
		return
	}
	err = ch.QueueBind(q.Name, userName.(string), "private", false, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
		c.Abort()
		return
	}
	client := &WS.Client{
		Close: make(chan struct{}),
		ID:    userID.(uint),
		Name:  userName.(string),
		MQCh:  ch,
		Queue: q,
		Hub:   WS.H,
		Conn:  conn,
		Send:  make(chan []byte, 256),
	}

	client.Hub.Register <- client

	// 每个连接两个 goroutine：读 + 写
	go client.WritePump()
	go client.ReadPump()
	fmt.Println("comsume begin")
	client.ConsumePrivate()
	fmt.Println("consume end")

}
