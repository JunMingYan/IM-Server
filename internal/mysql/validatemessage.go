package mysql

import (
	"server/internal/models"
)

func GetValidateMessage(vm models.ValidateMessageForm) (models.ValidateMessage, error) {
	var message models.ValidateMessage
	//SELECT * FROM validatemessage WHERE RoomID = 1 AND Status = 1 AND ValidateType = 1
	err := DB.Table("`validatemessage`").
		Select("*").
		Where("SenderID = ? AND ReceiverID = ? AND Status = ? AND ValidateType = ?", vm.SenderID, vm.ReceiverID, vm.Status, vm.ValidateType).
		Scan(&message).Error
	if err != nil {
		return models.ValidateMessage{}, err
	}

	return message, nil
}

func SaveValidateMessage(vm *models.ValidateMessage) error {
	err := DB.Create(vm).Error
	if err != nil {
		return err
	}
	return nil
}

func GetMyValidateMessageList(vm models.ValidateMessageListForm) ([]models.ValidateMessage, error) {
	var results []models.ValidateMessage
	//SELECT * FROM validatemessage WHERE RoomID = 1 AND Status = 1 AND ValidateType = 1
	err := DB.Table("`validatemessage`").
		Select("*").
		Where("ReceiverID = ?", vm.UserID).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func FindValidateMessage(message models.FriendVerificationForm) (models.ValidateMessage, error) {
	var result models.ValidateMessage

	err := DB.Table("`validatemessage`").
		Select("ID").
		Where("SenderID = ? AND ReceiverID = ? AND Status = ? AND ValidateType = ?",
			message.SenderID, message.ReceiverID, message.Status, message.ValidateType).
		Scan(&result).Error
	if err != nil {
		return models.ValidateMessage{}, err
	}

	return result, nil
}

func ModifyState(message models.FriendVerificationForm, status int) error {
	err := DB.Model(&models.ValidateMessage{}).
		Where("SenderID = ? AND ReceiverID = ? AND Status = ? AND ValidateType = ?",
			message.SenderID, message.ReceiverID, message.Status, message.ValidateType).
		Update("Status", status).Error
	if err != nil {
		return err
	}

	return nil
}
