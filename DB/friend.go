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
