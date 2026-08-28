package DB

import (
	"errors"
	"gorm.io/gorm/clause"
)

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
func InsertGroup(ownerID uint, ownerName, groupName string) error {
	var group = Group{
		GroupName: groupName,
		OwnerID:   ownerID,
		OwnerName: ownerName,
	}

	err := MySQL.Table("groups").Create(&group).Error
	return err
}
func InsertMember(memberID uint, groupName string) (err error) {
	groupID := QueryGroupID(groupName).ID
	if groupID == 0 {
		err = errors.New("NOT EXIST")

	} else {
		var member = GroupUser{
			GroupID: groupID,
			UserID:  memberID,
			Break:   false,
		}
		err = MySQL.Table("group_users").Create(&member).Error
	}

	return err
}
