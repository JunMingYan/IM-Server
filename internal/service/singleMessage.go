package service

import (
	"errors"
	"server/internal/models"
	"server/internal/mysql"
	"strconv"
	"strings"
)

func GetRecentSingleMessages(msg models.RecentSingleMessages) ([]*models.SingleMessage, error) {
	msg.Key = msg.RoomID

	results, err := mysql.GetRecentSingleMessages(msg.Key, msg.PageIndex, msg.PageSize)
	if err != nil {
		return nil, err
	}

	//for _, v := range results {
	//	list, err := StringToUint64(v.IsReadUser)
	//	if err != nil {
	//		logrus.Error(err)
	//		return nil, err
	//	}
	//	v.IsRead = list
	//}

	return results, nil
}

func SaveSingleMessages(msg models.SingleMessage) (models.SingleMessage, error) {
	return msg, mysql.SaveSingleMessages(msg)
}

func GetRoomID(roomID string) (uint64, error) {
	list := strings.Split(roomID, "-")
	if len(list) == 2 {
		id1, err := strconv.Atoi(list[0])
		if err != nil {
			return 0, err
		}
		ID1 := uint64(id1)
		id2, err := strconv.Atoi(list[1])
		if err != nil {
			return 0, err
		}
		ID2 := uint64(id2)
		if id1 > id2 {
			return ID1 - ID2, nil
		} else {
			return ID2 - ID1, nil
		}
	}
	return 0, errors.New("请重新输入RoomID")
}

func StringToUint64(str string) ([]uint64, error) {
	if str == "" {
		return nil, errors.New("str为空字符,请重新输入")
	}

	strList1 := strings.Split(str, "[")
	if len(strList1) == 0 {
		return nil, errors.New("strList长度为0,请重新输入")
	}

	strList2 := strings.Split(strList1[1], "]")
	if len(strList2) == 0 {
		return nil, errors.New("strList2长度为0,请重新输入")
	}

	strList := strings.Split(strList2[0], ",")
	if len(strList) == 0 {
		return nil, errors.New("strList长度为0,请重新输入")
	}

	var results []uint64
	for _, v := range strList {
		if v == "" {
			continue
		}

		result, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func LastFriendMsg(msg models.LastFriendMsg) (models.SingleMessage, error) {
	return mysql.LastFriendMsg(msg)
}

func IsRead(read models.IsRead) error {
	return mysql.IsRead(read)
}
