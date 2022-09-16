package mysql

import (
	models "server/internal/models"
	"time"
)

func RecentConversation(userID, goodFriendID uint64) ([]models.SimpleUser, error) {
	var goodFriends []models.SimpleUser

	/*	SELECT UserID,Photo,Signature,NickName,OnlineTime,LastLoginTime,UserName,CreateDate FROM `user`,`goodfriend`WHERE (
			(goodfriend.UserM = 8458548790951937 AND goodfriend.UserY = 8458566608355329)
		OR (goodfriend.UserY = 8458548790951937 AND goodfriend.UserM = 8458566608355329)
	) AND goodfriend.UserY = user.UserID*/
	err := DB.Table("`user`,`goodfriend`").
		Select("UserID,Photo,Signature,NickName,OnlineTime,LastLoginTime,UserName,CreateDate").
		Where("("+
			"(goodfriend.UserM = ? AND goodfriend.UserY = ?) "+
			"OR (goodfriend.UserY = ? AND goodfriend.UserM = ?)) "+
			"AND goodfriend.UserY = user.UserID", userID, goodFriendID, userID, goodFriendID).
		Scan(&goodFriends).Error
	if err != nil {
		return nil, err
	}

	return goodFriends, err
}

func MyFriendList(userID uint64) ([]*models.MyFriend, error) {
	var goodFriends []*models.MyFriend

	//SELECT CreateDate,NickName,Photo,Signature,UserID FROM goodfriend,user
	//WHERE goodfriend.UserM = 8458548790951937 AND user.UserID = goodfriend.UserY
	err := DB.Table("`user`,`goodfriend`").
		Select("CreateDate,NickName,Photo,Signature,UserID").
		Where("goodfriend.UserM = ? AND user.UserID = goodfriend.UserY", userID).
		Scan(&goodFriends).Error
	if err != nil {
		return nil, err
	}

	return goodFriends, nil
}

func FriendGroup(userID uint64) ([]models.FriendGroup, error) {
	var groups []models.FriendGroup

	//SELECT friendgroup.UserID,friendgroup.FriendID,friendgroupname.GroupName FROM friendgroup,friendgroupname
	//WHERE friendgroup.UserID = 8458548790951937 AND friendgroup.GroupID = friendgroupname.ID
	err := DB.Table("`friendgroup`,`friendgroupname`").
		Select("friendgroup.UserID,friendgroup.FriendID,friendgroupname.GroupName").
		Where("friendgroup.UserID = ? AND friendgroup.GroupID = friendgroupname.ID", userID).
		Scan(&groups).Error
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func FriendGroupName(userID uint64) ([]*models.FriendGroupName, error) {
	var nameList []*models.FriendGroupName

	//SELECT distinct friendgroupname.GroupName FROM `friendgroup`,`friendgroupname`
	//WHERE friendgroup.UserID = 8458548790951937
	err := DB.Table("`friendgroupname`").
		Select("distinct GroupName").
		Where("UserID = ?", userID). // AND friendgroup.GroupID = friendgroupname.ID
		Scan(&nameList).Error
	if err != nil {
		return nil, err
	}

	return nameList, nil
}

func FriendNotes(userID uint64) ([]*models.FriendNote, error) {
	var notes []*models.FriendNote

	//SELECT UserID,Notes,FriendID FROM friendnote where UserID = 8458548790951937
	err := DB.Table("`friendnote`").
		Select("UserID,Notes,FriendID").
		Where("UserID = ?", userID).
		Scan(&notes).Error
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func FindFriends(message models.FriendVerificationForm) ([]models.GoodFriend, error) {
	var results []models.GoodFriend

	err := DB.Table("`goodfriend`").
		Select("UserM,UserY,CreateDate").
		Where("(UserM = ? AND UserY = ?) OR (UserM = ? AND UserY = ?)", message.SenderID, message.ReceiverID, message.ReceiverID, message.SenderID).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func SaveFriend(me, friend uint64) error {
	goodFriend := &models.GoodFriend{
		UserM:      me,
		UserY:      friend,
		CreateDate: time.Now(),
	}
	err := DB.Create(goodFriend).Error
	if err != nil {
		return err
	}

	goodFriend = &models.GoodFriend{
		UserM:      friend,
		UserY:      me,
		CreateDate: time.Now(),
	}
	err = DB.Create(goodFriend).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateFriendNote(user models.UpdateFriendNote) error {
	err := DB.Model(&models.FriendNote{}).
		Where("UserID = ? AND FriendID = ?", user.UserID, user.FriendID).
		Update("Notes", user.FriendBeiZhuName).Error
	if err != nil {
		return err
	}

	return nil
}

func MoveFriendGroup(userID, friendID uint64, groupID int64, groupName string) error {
	err := DB.Model(models.FriendGroup{}).
		Where("UserID = ? AND FriendID = ?", userID, friendID).
		Updates(models.FriendGroup{GroupName: groupName, GroupID: groupID}).Error
	if err != nil {
		return err
	}

	return nil
}

func SaveFriendGroup(group *models.FriendGroupName) error {
	err := DB.Create(group).Error
	if err != nil {
		return err
	}
	return nil
}

func FindGroupNameByName(groupName string, userID uint64) (int64, error) {
	var result models.FriendGroupName

	err := DB.Table("`friendgroupname`").
		Select("ID").
		Where("GroupName = ? AND UserID = ?", groupName, userID).
		Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result.ID, nil
}

func ModifyToDefaultGroup(userID uint64, groupID, newID int64, groupName string) error {
	err := DB.Model(models.FriendGroup{}).
		Where("UserID = ? AND GroupID = ?", userID, groupID).
		Updates(models.FriendGroup{GroupName: groupName, GroupID: newID}).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteFriendGroup(deleteID int64) error {
	err := DB.Delete(&models.FriendGroupName{}, deleteID).Error
	if err != nil {
		return err
	}

	return nil
}

func FindFriendGroup(userID uint64, groupName string) ([]models.FriendGroup, error) {
	var result []models.FriendGroup

	err := DB.Table("`friendgroup`").
		Select("ID,GroupID").
		Where("GroupName = ? AND UserID = ?", groupName, userID).
		Scan(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func ModifyFriendGroupNameAndID(userID uint64, groupID int64, groupName string) error {
	err := DB.Model(models.FriendGroup{}).
		Where("UserID = ? AND GroupID = ?", userID, groupID).
		Updates(models.FriendGroup{GroupName: groupName, GroupID: groupID}).Error
	if err != nil {
		return err
	}

	return nil
}

func ModifyFriendGroupName(userID uint64, newName, oldName string) error {
	err := DB.Model(models.FriendGroupName{}).
		Where("UserID = ? AND GroupName = ?", userID, oldName).
		Updates(models.FriendGroup{GroupName: newName}).Error
	if err != nil {
		return err
	}

	return nil
}
