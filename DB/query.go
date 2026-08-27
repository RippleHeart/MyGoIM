package DB

import "errors"

func QueryFrdAll(userID uint) []User {
	var result []User
	MySQL.
		Raw("select b.name, b.id from users a,users b,user_users c where a.id=? and ((a.id=c.active_id and b.id =c.passive_id) or  (b.id=c.active_id and a.id =c.passive_id))", userID).
		Find(&result)
	return result
}
func QueryFrd(ID, frdID uint) User {
	var result User
	MySQL.Table("user_users").
		Where("((active_id = ? AND passive_id = ?) OR (active_id = ? AND passive_id = ?)) AND Break = ?",
			ID, frdID, frdID, ID, false).
		First(&result)
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
