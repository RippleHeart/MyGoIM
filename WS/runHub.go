package WS

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"mygoim/DB"

	"time"
)

var H *Hub

func InitHub() {
	H = NewHub()
	go H.Run()
}
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.Name] = client
			h.mu.Unlock()
			log.Printf("用户 %s 上线，当前在线: %d", client.Name, len(h.Clients))

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client.Name]; ok {
				delete(h.Clients, client.Name)
				close(client.Send)
			}
			h.mu.Unlock()
			log.Printf("用户 %s 下线，当前在线: %d", client.Name, len(h.Clients))

		case msg := <-h.Broadcast:
			data, _ := json.Marshal(msg)
			h.mu.RLock()
			for _, client := range h.Clients {
				select {
				case client.Send <- data:
				default:
					// 发送缓冲区满，关闭连接
					close(client.Send)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// ReadPump 持续读取客户端消息
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(512 * 1024) // 最大消息 512KB
	// 设置心跳检测（Pong 响应）
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("读取错误: %v", err)
			}
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			continue
		}
		msg.From = c.Name
		msg.ID = c.ID
		msg.Timestamp = time.Now().Unix()

		// 根据消息类型路由
		switch msg.Type {
		case "private":
			c.Hub.SendPrivate(msg)
		case "group":
			c.Hub.SendGroup(msg)
		case "system":
			c.Hub.Broadcast <- msg
		}
	}
}

// WritePump 持续向客户端写入消息
func (c *Client) WritePump() {
	ticker := time.NewTicker(54 * time.Second) // 心跳间隔
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (h *Hub) SendPrivate(msg Message) {
	var frd DB.User
	userTo, _ := DB.QueryUser(msg.To)
	if userTo.ID != 0 {
		frd = DB.QueryFrd(msg.ID, userTo.ID)
	}
	if frd.ID == 0 {
		h.mu.RLock()
		target, ok := h.Clients[msg.From]
		h.mu.RUnlock()
		data, _ := json.Marshal(NewSystemMsg("对方不是你的好友", msg.From))
		if ok {
			select {
			case target.Send <- data:
			default:
				close(target.Send)
			}
		}
	} else {
		data, _ := json.Marshal(msg)
		h.mu.RLock()
		target, ok := h.Clients[msg.To]
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
func (h *Hub) SendGroup(msg Message) {
	//todo  完善群发逻辑

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
