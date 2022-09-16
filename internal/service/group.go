package service

import (
	models2 "server/internal/models"
	mysql2 "server/internal/mysql"
)

func MyGroupList(userName string) ([]*models2.MyGroup, error) {
	results, err := mysql2.MyGroupList(userName)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		user, err := mysql2.FindSimpleUser(userName)
		if err != nil {
			return nil, err
		}
		v.UserInfo = user

		group, err := mysql2.FindGroupByGroupID(v.GroupID)
		if err != nil {
			return nil, err
		}
		v.GroupInfo = group
	}
	return results, nil
}

func RecentGroupList(group models2.RecentGroupForm) ([]*models2.MyGroup, error) {
	var results []*models2.MyGroup
	for _, v := range group.GroupIDList {
		recentGroup, err := mysql2.RecentGroup(group.UserID, v)
		if err != nil {
			return nil, err
		}

		group, err := mysql2.FindGroupByGroupID(v)
		if err != nil {
			return nil, err
		}
		recentGroup.GroupInfo = group
		results = append(results, recentGroup)
	}

	return results, nil
}

func RecentGroupMessage(message models2.RecentGroupMessage) ([]*models2.GroupMessage, error) {
	return mysql2.RecentGroupMessage(message)
}

func GroupInfo(group models2.GroupInfoForm) (*models2.Group, []*models2.User, error) {
	result, err := mysql2.GroupInfo(group)
	if err != nil {
		return nil, nil, err
	}

	results, err := mysql2.GroupUserList(group)
	if err != nil {
		return nil, nil, err
	}

	return result, results, nil
}

func SaveGroupMessage(message models2.GroupMessage) error {
	return mysql2.SaveGroupMessage(message)
}

func GroupUserList(groupID uint64) ([]*models2.User, error) {
	group := models2.GroupInfoForm{GroupID: groupID}
	results, err := mysql2.GroupUserList(group)
	if err != nil {
		return nil, err
	}

	return results, nil
}
