package Chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"mygoim/DB"
)

func InitChat(userName string, userID uint, conn *websocket.Conn) error {

	//创建AMQP信道
	ch, err := NewChannel()
	if err != nil {
		return err
	}
	//声明收消息的队列
	q, err := ch.QueueDeclare(userName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	//将队列绑定到交换机上
	err = ch.QueueBind(q.Name, userName, "private", false, nil)
	if err != nil {
		return err
	}
	groups := DB.QueryMyGroup(userID)
	for _, group := range groups {
		if group.GroupName != "" {
			err = ch.QueueBind(q.Name, group.GroupName, "group", false, nil)
			if err != nil {
				return err
			}
		}
	}

	//注册WS客户端连接
	client := &Client{
		Close: make(chan struct{}),
		ID:    userID,
		Name:  userName,
		MQCh:  ch,
		Queue: q,
		Hub:   H,
		Conn:  conn,
		Send:  make(chan []byte, 256),
	}
	client.Hub.Register <- client
	//Redis进行缓存上线记录
	DB.SetOnline(userName)
	fmt.Println(DB.CheckOnline(userName))
	fmt.Println(DB.CheckOnline("userName"))
	fmt.Println(DB.CheckOnline("123"))
	//开辟 读+写goroutine、消费消息的goroutine
	go client.WritePump()
	go client.ReadPump()
	go client.ConsumeMyQueue()
	return nil
}
