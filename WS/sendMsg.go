package WS

import (
	"encoding/json"
	"fmt"
	"mygoim/DB"
)

func (c *Client) SendPrivate(msg Message) {
	var frd DB.User
	//判断是否为好友
	userTo, _ := DB.QueryUser(msg.To)
	if userTo.ID != 0 {
		frd = DB.QueryFrd(msg.ID, userTo.ID)
	}
	if frd.ID == 0 { //对象不是好友返回系统消息
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
	//对象是好友
	//publish到私聊交换机
	fmt.Println("进入publish")
	err := c.PublishPrivate(msg)
	if err != nil { //发送失败
		fmt.Println("publish失败")
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

}
func (h *Hub) SendGroup(msg Message) {

	usersTo := DB.QueryMemberAll(msg.To)
	for _, user := range usersTo {
		data, _ := json.Marshal(msg)
		h.mu.RLock()
		target, ok := h.Clients[user.Name]
		h.mu.RUnlock()
		if ok {
			select {
			case target.Send <- data:
			default:
				close(target.Send)
			}
		}
	}
}
