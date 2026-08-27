package DB

import "gorm.io/gorm/clause"

func InsertFrd(ID, frdID uint) error {
	//删除标记恢复机制
	err := MySQL.Table("user_users").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "active_id"}, {Name: "passive_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"break", "updated_at"}),
	}).Create(&UserUser{
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
