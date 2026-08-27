package DB

func InsertFrd(ID, frdID uint) error {
	err := MySQL.Table("user_users").Create(&UserUser{
		ActiveID:  ID,
		PassiveID: frdID,
		Break:     false,
	}).Error
	return err
}

func InsertUser(name string, pwd string) error {
	var user = User{
		Name:     name,
		Password: pwd,
	}
	err := MySQL.Table("users").Create(&user).Error
	return err

}
