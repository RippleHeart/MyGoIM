package DB

import "errors"

func InsertUser(name string, pwd string) error {
	var user = User{
		Name:     name,
		Password: pwd,
	}
	err := MySQL.Table("users").Create(&user).Error
	return err

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
