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
func QueryGroupID(input string) Group {
	var result Group
	MySQL.Table("groups").Select("ID").Where("group_name=?", input).First(&result)
	return result
}
func QueryMemberAll(groupName string) []User {
	var result []User

	groupID := QueryGroupID(groupName).ID
	if groupID == 0 {
		return result
	}
	MySQL.
		Raw("select a.name, a.id from users a,`groups` b,group_users c where c.break=0 and b.id=? and b.id=c.group_id and a.id=c.user_id ",
			groupID).
		Find(&result)
	return result
}

func QueryMember(userName, groupName string) User {
	var result User

	groupID := QueryGroupID(groupName).ID
	if groupID == 0 {
		return result
	}
	MySQL.
		Raw("select a.name, a.id from users a,group_users c where c.break=0 and c.group_id=? and a.id=c.user_id and a.name=?",
			groupID, userName).
		Find(&result)
	return result
}
func QueryMyGroup(input uint) []Group {
	var result []Group

	MySQL.
		Raw("select b.group_name, b.id from users a,`groups` b,group_users c where c.break=0 and a.id=? and b.id=c.group_id and a.id=c.user_id ",
			input).
		Find(&result)
	return result
}
