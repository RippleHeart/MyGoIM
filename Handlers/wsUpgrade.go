package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"mygoim/DB"
	"mygoim/WS"
	"net/http"
)

func WSUpgrade(c *gin.Context) {
	//升级HTTP为WS
	conn, err := WS.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "WS UP ERROR!"})
		c.Abort()
		return
	}

	userID, _ := c.Get("ID")
	userName, _ := c.Get("name")
	//创建AMQP信道
	ch, err := WS.NewChannel()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
		c.Abort()
		return
	}
	//声明收消息的队列
	q, err := ch.QueueDeclare(userName.(string), true, false, false, false, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
		c.Abort()
		return
	}
	//将队列绑定到交换机上
	err = ch.QueueBind(q.Name, userName.(string), "private", false, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
		c.Abort()
		return
	}
	groups := DB.QueryMyGroup(userID.(uint))
	for _, group := range groups {
		if group.GroupName != "" {
			err = ch.QueueBind(q.Name, group.GroupName, "group", false, nil)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"code": "102", "msg": "Try Again!"})
				c.Abort()
				return
			}
		}
	}

	//注册WS客户端连接
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

	//开辟 读+写goroutine、消费消息的goroutine
	go client.WritePump()
	go client.ReadPump()
	go client.ConsumePrivate()

}
