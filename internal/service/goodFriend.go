package service

import (
	"errors"
	"server/internal/models"
	"server/internal/mysql"
)

const DefaultLevel = 1

const DefaultFriend = "我的好友"

// 每天在线一小时升一级，不足一小时会进行累计，累计满一小时升一级
func computedLevel(onlineTime int64) int {
	return DefaultLevel
}

func MyFriendList(userID uint64) ([]*models.MyFriend, error) {
	results, err := mysql.MyFriendList(userID)
	if err != nil {
		return nil, err
	}

	for _, v := range results {
		v.Level = computedLevel(v.OnlineTime)
		if userID > v.UserID {
			v.RoomID = userID - v.UserID
		} else {
			v.RoomID = v.UserID - userID
		}
	}

	return results, err
}

func UpdateFriendNote(user models.UpdateFriendNote) (map[uint64]string, error) {
	err := mysql.UpdateFriendNote(user)
	if err != nil {
		return nil, err
	}

	notes, err := mysql.FriendNotes(user.UserID)
	results := make(map[uint64]string)
	for _, v := range notes {
		results[v.FriendID] = v.Notes
	}

	return results, nil
}

/*
	修改分组
	1.默认分组不能修改
	2.通过UserID和FriendGroupName查找FriendGroupName的ID
	3.通过UserID和FriendID修改FriendGroup的GroupName和GroupID
	4.返回最新好友分组的数据
*/
func MoveFriendGroup(user models.UpdateFriendGroup) (map[string][]uint64, error) {
	// 1.默认分组不能修改
	if DefaultFriend == user.FriendGroup {
		return nil, errors.New("默认分组不能修改")
	}
	// 2.
	ID, err := mysql.FindGroupNameByName(user.FriendGroup, user.UserID)
	if err != nil {
		return nil, err
	}

	err = mysql.MoveFriendGroup(user.UserID, user.FriendID, ID, user.FriendGroup)
	if err != nil {
		return nil, err
	}

	results, err := findFriendGroup(user.UserID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

/*
	新增分组
	1.检查新增分组是否已经分组
	2.返回最新好友分组数据
*/
func SaveFriendGroup(group models.AddFriendGroup) (map[string][]uint64, error) {

	ID, err := mysql.FindGroupNameByName(group.FenZuName, group.UserID)
	if err != nil {
		return nil, err
	}
	if ID > 0 {
		return nil, errors.New("分组已存在,请重新输入分组")
	}

	newGroup := &models.FriendGroupName{
		UserID:    group.UserID,
		GroupName: group.FenZuName,
	}

	err = mysql.SaveFriendGroup(newGroup)
	if err != nil {
		return nil, err
	}

	results, err := findFriendGroup(group.UserID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

/*
	删除分组
	1.默认分组不能删除
	2.通过UserID和GroupName查找FriendGroupName的ID,查找默认分组的ID
	3.把要删除的好友分组中的好友移动到默认好友分组
	4.删除掉需要删除的好友分组
	5.返回最新好友分组数据
*/
func DeleteFriendGroup(group models.DeleteFriendGroup) (map[string][]uint64, error) {
	// 1.默认分组不能删除
	if DefaultFriend == group.FenZuName {
		return nil, errors.New("默认分组不能删除")
	}
	// 2.把要删除的分组中的好友先修改到默认分组
	defaultID, err := mysql.FindGroupNameByName(DefaultFriend, group.UserID)
	if err != nil {
		return nil, err
	}
	deleteID, err := mysql.FindGroupNameByName(group.FenZuName, group.UserID)
	if err != nil {
		return nil, err
	}
	err = mysql.ModifyToDefaultGroup(group.UserID, deleteID, defaultID, DefaultFriend)
	if err != nil {
		return nil, err
	}
	// 3.删除分组
	//deleteGroup := &models.FriendGroupName{ID: defaultID}
	err = mysql.DeleteFriendGroup(deleteID)
	if err != nil {
		return nil, err
	}

	results, err := findFriendGroup(group.UserID)
	if err != nil {
		return nil, err
	}

	return results, nil
}

/*
	修改分组名称
	1.默认分组不能修改
	2-1.该分组不存在好友
	3.直接修改friendGroupName
	4.返回最新的好友分组
	2-2.该分组存在好友
	3.检查新的分组名称是否存在
	4.检查旧的分组名称是否存在
	5.修改分组名称
	6.返回最新的好友分组
*/
func ModifyFriendGroupName(user models.UpdateFriendGroupName) (map[string][]uint64, error) {
	// 1.默认分组不能修改
	if DefaultFriend == user.OldFenZu {
		return nil, errors.New("默认分组不能修改")
	}

	list, err := mysql.FindFriendGroup(user.UserID, user.OldFenZu)
	if err != nil {
		return nil, err
	}

	if len(list) != 0 { // 分组存在好友
		ID, err := mysql.FindGroupNameByName(user.NewFenZu, user.UserID)
		if err != nil {
			return nil, err
		}
		if ID != 0 {
			return nil, errors.New("分组名称已存在")
		}

		ID, err = mysql.FindGroupNameByName(user.OldFenZu, user.UserID)
		if err != nil {
			return nil, err
		}

		err = mysql.ModifyFriendGroupNameAndID(user.UserID, ID, user.NewFenZu)
		if err != nil {
			return nil, err
		}

		err = mysql.ModifyFriendGroupName(user.UserID, user.NewFenZu, user.OldFenZu)
		if err != nil {
			return nil, err
		}
	}

	err = mysql.ModifyFriendGroupName(user.UserID, user.NewFenZu, user.OldFenZu)
	if err != nil {
		return nil, err
	}

	results, err := findFriendGroup(user.UserID)
	if err != nil {
		return nil, err
	}

	return results, nil
}
