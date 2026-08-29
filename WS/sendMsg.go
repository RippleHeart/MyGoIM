package WS

import (
	"encoding/json"
	"mygoim/DB"
)

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
