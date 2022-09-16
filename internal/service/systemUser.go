package service

import (
	"server/internal/models"
	"server/internal/mysql"
)

func SystemUserList() ([]models.SystemUser, error) {

	return mysql.SystemUserList()
}

func SaveFeedBack(back models.FeedBack) error {
	return mysql.SaveFeeBack(back)
}
