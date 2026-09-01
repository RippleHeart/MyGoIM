package Chat

import (
	"encoding/json"
	"log"
	"mygoim/DB"
	"time"
)

func (c *Client) SendPrivate(msg Message) {
	var frd DB.User
	// 判断是否为好友
	userTo, _ := DB.QueryUser(msg.To)
	if userTo.ID != 0 {
		frd = DB.QueryFrd(msg.ID, userTo.ID)
	}
	// 对方不是好友返回系统消息
	if frd.ID == 0 {
		c.Hub.mu.RLock()
		target, ok := c.Hub.Clients[msg.From]
		c.Hub.mu.RUnlock()
		data, _ := json.Marshal(NewSystemMsg("对方不是你的好友", msg.From))
		if ok {
			select {
			case target.Send <- data:
			default:
				close(target.Send)
			}
		}
		return
	}
	// Publish到私聊交换机
	err := c.PublishPrivate(msg)
	log.Println("published")
	// Publish失败通知
	if err != nil {
		c.Hub.mu.RLock()
		target, ok := c.Hub.Clients[msg.From]
		c.Hub.mu.RUnlock()
		data, _ := json.Marshal(NewSystemMsg("发送失败", msg.From))
		if ok {
			select {
			case target.Send <- data:
			default:
				close(target.Send)
			}
		}
		return
	}
	//持久化到MySQL
	data, _ := json.Marshal(msg)
	err = DB.InsertMsg(msg.ID, frd.ID, msg.Type, data)
	if err != nil {
		//todo 持久化失败原因及处理
		return
	}
}

func (c *Client) SendGroup(msg Message) {
	// 判断发送者是否为群成员
	if DB.QueryMember(msg.From, msg.To).ID == 0 {
		c.Hub.mu.RLock()
		target, ok := c.Hub.Clients[msg.From]
		c.Hub.mu.RUnlock()
		data, _ := json.Marshal(NewSystemMsg("你不是该群成员", msg.From))
		if ok {
			select {
			case target.Send <- data:
			default:
				close(target.Send)
			}
		}
		return
	}
	// Publish到群聊交换机
	err := c.PublishGroup(msg)
	// Publish失败通知
	if err != nil {
		c.Hub.mu.RLock()
		target, ok := c.Hub.Clients[msg.From]
		c.Hub.mu.RUnlock()
		data, _ := json.Marshal(NewSystemMsg("发送失败", msg.From))
		if ok {
			select {
			case target.Send <- data:
			default:
				close(target.Send)
			}
		}
		return
	}
	//持久化到数据库
	group := DB.QueryGroupID(msg.To)
	data, _ := json.Marshal(msg)
	err = DB.InsertMsg(msg.ID, group.ID, msg.Type, data)
	if err != nil {
		//todo 持久化失败原因及处理
		return
	}
}

func NewSystemMsg(content string, to string) *Message {
	return &Message{
		Type:      "system",
		From:      "system",
		ID:        0,
		To:        to,
		Content:   content,
		Timestamp: time.Now().Unix(),
	}
}
