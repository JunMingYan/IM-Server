package models

import "time"

const (
	StatusUntreated = iota
	StatusAgree
	StatusRefuse
)

const (
	ValidateTypeFriend = iota
	ValidateTypeGroup
)

type ValidateMessage struct {
	ID              int       `gorm:"column:ID;primaryKey;" json:"-" form:"-"`
	RoomID          string    `gorm:"-" json:"-" form:"-"` // column:RoomID;type:string;not null;size:32 roomId
	SenderID        uint64    `gorm:"column:SenderID;type:string;not null;size:32" json:"SenderID" form:"SenderID"`
	SenderName      string    `gorm:"column:SenderName;type:string;not null;size:32" json:"senderName"form:"senderName"`
	SenderNickname  string    `gorm:"column:SenderNickname;type:string;not null;size:32" json:"senderNickname"form:"senderNickname"`    // 发送者昵称
	SenderAvatar    string    `gorm:"column:SenderAvatar;type:string;not null;size:32" json:"senderAvatar"form:"senderAvatar"`          // 发送者头像
	ReceiverID      uint64    `gorm:"column:ReceiverID;type:string;not null;size:32" json:"receiverID"form:"receiverID"`                // 接收者ID
	Time            time.Time `gorm:"column:Time;type:time;not null;size:32" json:"time"form:"time"`                                    // 消息发送时间
	AdditionMessage string    `gorm:"column:AdditionMessage;type:string;not null;size:32" json:"additionMessage"form:"additionMessage"` // 附加消息
	Status          int       `gorm:"column:Status;type:int;not null;size:32" json:"status"form:"status"`                               // 0/1/2，未处理/同意/不同意
	ValidateType    int       `gorm:"column:ValidateType;type:int;not null;size:32" json:"validateType"form:"validateType"`             // 0/1, 好友/群聊
	GroupID         uint64    `gorm:"-" json:"-"form:"-"`                                                                               //column:GroupID;type:string;not null;size:32
}

func (ValidateMessage) TableName() string {
	return "validateMessage"
}

type ValidateMessageForm struct {
	SenderID     int64 `json:"senderID" form:"senderID" `
	ReceiverID   int64 `json:"receiverID" form:"receiverID" `
	Status       int   `json:"status" form:"status" `
	ValidateType int   `json:"validateType" form:"validateType" `
}

type ValidateMessageListForm struct {
	UserID uint64 `json:"userID"  form:"userID"`
}

type FriendVerificationForm struct {
	SenderID        uint64    `json:"senderID"`
	AdditionMessage string    `json:"additionMessage"`
	ReceiverID      uint64    `json:"receiverID"`
	SenderAvatar    string    `json:"senderAvatar"`
	SenderName      string    `json:"senderName"`
	SenderNickname  string    `json:"senderNickname"`
	Status          int       `json:"status"`
	Time            time.Time `json:"time"`
	ValidateType    int       `json:"validateType"`
}
