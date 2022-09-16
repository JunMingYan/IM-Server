package mysql

import (
	"server/internal/models"
)

func GetRecentSingleMessages(key uint64, pageIndex, pageSize int) ([]*models.SingleMessage, error) {
	var results []*models.SingleMessage
	if pageIndex-1 > 0 {
		pageIndex = pageIndex - 1
	}

	//SELECT * FROM singlemessage WHERE RoomID = '8458548790951937-8458566608355329' ORDER BY SendTime DESC LIMIT 0,10
	err := DB.Table("singlemessage").
		Select("*").
		Where("`RoomID` = ? ORDER BY SendTime DESC LIMIT ?,?", key, pageIndex*pageSize, pageSize).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func SaveSingleMessages(msg models.SingleMessage) error {
	err := DB.Create(&msg).Error
	if err != nil {
		return err
	}
	return nil
}

//SELECT * FROM singlemessage WHERE `key` = 17817403392 ORDER BY SendTime ASC LIMIT 20,10

func LastFriendMsg(msg models.LastFriendMsg) (models.SingleMessage, error) {
	var result models.SingleMessage

	//SELECT * FROM singlemessage WHERE RoomID = "8458548790951937-8458566608355329" ORDER BY SendTime DESC LIMIT 0,1
	err := DB.Table("singlemessage").
		Select("*").
		Where("RoomID = ? ORDER BY SendTime DESC LIMIT 0,1", msg.RoomID).
		Scan(&result).Error
	if err != nil {
		return models.SingleMessage{}, err
	}
	return result, nil
}

func IsRead(read models.IsRead) error {
	//UPDATE singlemessage SET IsRead = 1 WHERE IsRead = 0 AND RoomID = 17817403392 AND ReceiverID = 8458566608355329

	err := DB.Model(&models.SingleMessage{}).
		Where("IsRead = ? AND RoomID = ? AND ReceiverID = ?", 0, read.RoomID, read.UserID).
		Update("IsRead", 1).Error
	if err != nil {
		return err
	}
	return nil
}
