package DB

func InsertGroup(ownerID uint, ownerName, groupName string) error {
	var group = Group{
		GroupName: groupName,
		OwnerID:   ownerID,
		OwnerName: ownerName,
	}

	err := MySQL.Table("groups").Create(&group).Error
	return err
}
func QueryGroupID(input string) Group {
	var result Group
	MySQL.Table("groups").Select("ID").Where("group_name=?", input).First(&result)
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
