package mysql

import (
	"server/internal/models"
)

func SaveSystemUser(sys models.SystemUser) error {
	err := DB.Create(sys).Error
	if err != nil {
		return err
	}
	return nil
}

func SystemUserList() ([]models.SystemUser, error) {
	var systemUsers []models.SystemUser
	result := DB.Find(&systemUsers)
	if result.Error != nil {
		return nil, result.Error
	}

	return systemUsers, nil
}

func SaveFeeBack(back models.FeedBack) error {
	err := DB.Create(&back).Error
	if err != nil {
		return err
	}
	return nil
}
