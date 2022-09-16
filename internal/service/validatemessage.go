package service

import (
	"errors"
	"server/internal/models"
	"server/internal/mysql"
	"time"
)

func GetValidateMessage(vm models.ValidateMessageForm) (models.ValidateMessage, error) {
	return mysql.GetValidateMessage(vm)
}

func SendValidateMessage(vm *models.ValidateMessage) error {
	vm.Time = time.Now()
	return mysql.SaveValidateMessage(vm)
}

func GetMyValidateMessageList(vm models.ValidateMessageListForm) ([]models.ValidateMessage, error) {
	return mysql.GetMyValidateMessageList(vm)
}

func Reject(message models.FriendVerificationForm) error {
	friend, err := mysql.FindValidateMessage(message)
	if err != nil {
		return err
	}
	if friend.ID <= 0 {
		return errors.New("好友验证消息不存在")
	}

	err = mysql.ModifyState(message, models.StatusRefuse)
	if err != nil {
		return err
	}

	return nil
}

func Agree(message models.FriendVerificationForm) error {
	// 2.查询是否已经是好友关系，只要不是好友关系就可以添加为好友
	friend, err := mysql.FindValidateMessage(message)
	if err != nil {
		return err
	}
	if friend.ID <= 0 {
		return errors.New("好友验证消息不存在")
	}
	results, err := mysql.FindFriends(message)
	if err != nil {
		return err
	}
	if len(results) == 2 {
		return errors.New("不能重复添加好友")
	}
	// 3.添加好友关系
	err = mysql.ModifyState(message, models.StatusAgree)
	if err != nil {
		return err
	}
	err = mysql.SaveFriend(message.SenderID, message.ReceiverID)
	if err != nil {
		return err
	}
	// TODO 4.订阅消息队列

	return nil
}
