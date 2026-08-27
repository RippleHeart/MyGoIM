package WS

import (
	"github.com/gorilla/websocket"
	"mygoim/Conf"
	"net/http"
	"sync"
)

type Message struct {
	Type      string `json:"type"` //  "private" | "system" | "group"
	From      string `json:"from"` // 发送者
	ID        uint   `json:"id"`
	To        string `json:"to"`        // 接收者
	Content   string `json:"content"`   // 消息内容
	Timestamp int64  `json:"timestamp"` // 时间戳
}

// Client 代表一个 WebSocket 连接
type Client struct {
	ID   uint
	Name string
	Conn *websocket.Conn
	Send chan []byte // 待发送的消息队列
	Hub  *Hub
}

// Hub 管理所有客户端连接
type Hub struct {
	Clients    map[string]*Client // userID -> Client
	Register   chan *Client       // 新连接注册
	Unregister chan *Client       // 连接断开注销
	Broadcast  chan Message       // 广播消息
	mu         sync.RWMutex       // 保护 Clients map
}

var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  Conf.Conf.W.ReadBufferSize,
	WriteBufferSize: Conf.Conf.W.WriteBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true // 生产环境应校验域名
	},
}
