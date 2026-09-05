package Service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mygoim/DB"
	"net/http"
	"strconv"
	"time"
)

type BasicInfo struct {
	UserNum      uint `json:"register_num"`
	OnlineNum    uint `json:"online_num"`
	MsgToday     uint `json:"msg_today"`
	FresherToday uint `json:"fresher_today"`
}

func AdminInit() {
	go func() {
		for {
			<-time.Tick(time.Hour * 24)
			DB.RDB.Set(context.TODO(), "messageNum", 0, 0)
		}
	}()
}
func GetBasicInfo(c *gin.Context) {
	var basicInfo BasicInfo
	basicInfo.UserNum = DB.UserNum()
	basicInfo.OnlineNum = DB.OnlineNum()
	basicInfo.MsgToday = DB.GetMsgNum()
	basicInfo.FresherToday = DB.GetFresherToday()
	c.JSON(http.StatusOK, gin.H{"code": "0", "basicinfo": basicInfo})
}
func QueryMessage(c *gin.Context) {
	userid := c.Query("id")
	msgType := c.Query("type")
	parseUint, err := strconv.ParseUint(userid, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "1001", "error": "ID can't parse"})
	}
	msgs, err := DB.QueryUserMsg(uint(parseUint), msgType)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "1001", "error": "Type can't search"})
	}
	// TODO !!! 解码问题
	for i, msg := range msgs {
		decodedBytes, _ := base64.StdEncoding.DecodeString(string(msg.Content))
		_ = json.Unmarshal(decodedBytes, &msgs[i].Content)
	}
	c.JSON(http.StatusOK, gin.H{"code": "1001", "messages": msgs})
}
