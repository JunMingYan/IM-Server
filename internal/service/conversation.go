package service

import (
	"github.com/sirupsen/logrus"
	models "server/internal/models"
	"server/internal/mysql"
)

func ConversationList(c models.ConversationForm) ([]models.SingleRecentConversation, error) {
	var results []models.SingleRecentConversation
	var item models.SingleRecentConversation

	for _, v := range c.RecentFriendIds {
		userList, err := mysql.RecentConversation(c.UserID, v)
		if err != nil {
			logrus.Fatal(err)
			return nil, err
		}

		if c.UserID == userList[0].UserID {
			item.UserID = c.UserID
			item.UserY = userList[1]
			item.UserM = userList[0]
			item.CreateDate = userList[0].CreateDate
		} else {
			item.UserID = c.UserID
			item.UserY = userList[0]
			item.UserM = userList[1]
			item.CreateDate = userList[1].CreateDate
		}

		if c.UserID > userList[0].UserID {
			item.RoomID = c.UserID - userList[0].UserID
		} else {
			item.RoomID = userList[0].UserID - c.UserID
		}

		results = append(results, item)
	}

	return results, nil
}
