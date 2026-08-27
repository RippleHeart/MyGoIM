package DB

import "errors"

func QueryFrdAll(userID uint) []User {
	var result []User
	MySQL.
		Raw("select b.name, b.id from users a,users b,user_users c where c.break=0 and a.id=? and ((a.id=c.active_id and b.id =c.passive_id) or  (b.id=c.active_id and a.id =c.passive_id))", userID).
		Find(&result)
	return result
}
func QueryFrd(ID, frdID uint) User {
	var result User
	MySQL.
		Raw("select a.name, a.id from users a,user_users c where c.break=0 and a.id=? and ((c.active_id=? and c.passive_id=?) or  (c.active_id=? and c.passive_id=?))", frdID, ID, frdID, frdID, ID).
		Find(&result)
	return result
}
func QueryUser(input any) (User, error) {
	var result User
	switch input.(type) {
	case uint:
		MySQL.Table("users").Select("*").Where("id=?", input).First(&result)
	case string:
		MySQL.Table("users").Select("*").Where("name=?", input).First(&result)
	default:
		return User{}, errors.New("input type error")
	}
	return result, nil
}
func QueryPwd(input string) User {
	var result User
	MySQL.Table("users").Select("password").Where("name=?", input).First(&result)
	return result
}
