package DB

import "errors"

func InsertMsg(from, to uint, MsgType string, Content []byte) (err error) {
	var msgs = PrivateMessage{
		ToID:    to,
		FromID:  from,
		Content: Content,
	}

	switch MsgType {
	case "private":

		err = MySQL.Table("private_messages").Create(&msgs).Error

	case "group":

		err = MySQL.Table("group_messages").Create(&msgs).Error

	default:
		err = errors.New("message type error")
	}
	return err
}

func QueryMsgAll(ID uint, IDType string) (err error) {
	var msgs PrivateMessage
	switch IDType {
	case "private":
		err = MySQL.Table("private_messages").Select("content").Where("from_id=? or to_id=?", ID, ID).Find(&msgs).Error

	case "group":
		err = MySQL.Table("group_messages").Select("content").Where("to_id=?", ID).Find(&msgs).Error
	default:
		err = errors.New("message type error")
	}
	return err
}
