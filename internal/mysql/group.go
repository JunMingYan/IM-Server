package mysql

import (
	models2 "server/internal/models"
)

func CreateGroup() {

}

func MyGroupList(userName string) ([]*models2.MyGroup, error) { // TODO 重构该方法
	var groups []*models2.MyGroup

	//SELECT g.ID,u.UserID,u.UserName,u.Manage,u.Holder,u.Card,u.Time,g.GroupID FROM `group` AS g,groupuser AS u
	//WHERE u.UserName = "JULIA" AND u.GroupID = g.GID
	err := DB.Table("`group`,`groupuser`").
		Select("group.ID,group.GroupID,groupuser.UserID,groupuser.UserName,groupuser.Role,groupuser.Card,groupuser.Time").
		Where("groupuser.UserName = ? AND groupuser.GroupID = group.GroupID", userName).
		Scan(&groups).Error
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func FindGroupByGroupID(groupID int64) (group models2.Group, err error) {
	err = DB.Where("GroupID = ?", groupID).First(&group).Error
	if err != nil {
		return models2.Group{}, err
	}
	return
}

func RecentGroupTest(UserID, groupID int64) (*models2.MyGroup, error) { // TODO 重构该方法 配置好MyGroup中gorm和mysql的字段映射
	var group *models2.MyGroup
	//SELECT
	//`group`.GroupID,groupuser.UserID,groupuser.UserName,groupuser.Role,groupuser.Card,groupuser.Time,`group`.GroupID,`group`.Title,
	//	`group`.`Desc`,`group`.Img,`group`.`Code`,`group`.UserNum,
	//	`group`.CreateDate,HolderName,`group`.HolderUserID,
	//	`user`.UserID,`user`.Photo,`user`.Signature,`user`.NickName,
	//	`user`.OnlineTime,`user`.LastLoginTime,`user`.UserName,groupuser.Time
	//FROM `group`,groupuser,`user`
	//WHERE `group`.GroupID IN (111,222,333)
	//AND groupuser.UserID = 8458548790951937
	//AND `group`.GroupID = groupuser.GroupID
	//AND groupuser.UserID = `user`.UserID

	err := DB.Table("`group`,`groupuser`").
		Select("group.ID,group.GroupID,groupuser.UserID,groupuser.UserName,groupuser.Role,groupuser.Card,groupuser.Time").
		Where("groupuser.GroupID = ? AND groupuser.UserID = ? AND group.GroupID = groupuser.GroupID", groupID, UserID).
		Scan(&group).Error
	if err != nil {
		return nil, err
	}

	return group, nil
}

func RecentGroup(UserID, groupID int64) (*models2.MyGroup, error) {
	var group *models2.MyGroup

	//SELECT * FROM `group` AS g,groupuser AS u
	//WHERE u.GroupID = 111 AND u.UserID = 8458548790951937 AND g.GroupID = u.GroupID
	err := DB.Table("`group`,`groupuser`").
		Select("group.ID,group.GroupID,groupuser.UserID,groupuser.UserName,groupuser.Role,groupuser.Card,"+
			"groupuser.Time").
		Where("groupuser.GroupID = ? AND groupuser.UserID = ? AND group.GroupID = groupuser.GroupID", groupID, UserID).
		Scan(&group).Error
	if err != nil {
		return nil, err
	}

	return group, nil
}

func RecentGroupMessage(message models2.RecentGroupMessage) ([]*models2.GroupMessage, error) {
	var results []*models2.GroupMessage
	if message.PageIndex-1 > 0 {
		message.PageIndex = message.PageIndex - 1
	}
	//SELECT * FROM groupMessage
	//WHERE RoomID = 17817403392 ORDER BY SendTime ASC LIMIT 20,10
	err := DB.Table("`groupMessage`").
		Select("*").
		Where("RoomID = ? ORDER BY SendTime ASC LIMIT ?,?",
			message.RoomID, message.PageIndex*message.PageSize, message.PageSize).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GroupInfo(group models2.GroupInfoForm) (*models2.Group, error) {
	var result *models2.Group

	//SELECT * FROM `group` WHERE GroupID = 111
	err := DB.Table("`group`").
		Select("*").
		Where("GroupID = ?", group.GroupID).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GroupUserList(group models2.GroupInfoForm) ([]*models2.User, error) {
	var results []*models2.User

	//SELECT * FROM groupuser WHEREUserID IN (SELECT UserID FROM groupuser WHERE GroupID = 111)
	err := DB.Table("`user`").
		Select("*").
		Where("UserID IN (SELECT UserID FROM groupuser WHERE GroupID = ?)", group.GroupID).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func SaveGroupMessage(message models2.GroupMessage) error {
	err := DB.Create(&message).Error
	if err != nil {
		return err
	}
	return nil
}
