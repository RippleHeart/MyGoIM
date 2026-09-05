package DB

func InsertUserInfo(id uint) error {
	var userInfo = UserInfo{
		UserID: id,
	}
	err := MySQL.Table("user_infos").Create(&userInfo).Error
	return err
}
