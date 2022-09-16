package models

import (
	"time"
)

//type Message struct {
//	ID          int       `db:"F_ID"`
//	MsgID       uint64    `db:"F_MsgID"`
//	MsgType     int       `json:"msgType" form:"msgType" binding:"required"`
//	SendID      uint64    `json:"sendID" form:"sendID" db:"F_SendID" binding:"required"`
//	RecipientID uint64    `json:"recipientID" form:"recipientID" db:"F_RecipientID" binding:"required"`
//	Content     string    `json:"content" form:"content" db:"F_Content" binding:"required"`
//	SendTime    time.Time `db:"F_SendTime"`
//}

type Message struct {
	ID               uint64    `json:"-"`
	MsgID            uint64    `json:"-"`
	ConversationType string    `json:"conversationType"`
	IsReadUser       []string  `json:"-"`
	Message          string    `json:"message"`
	MessageType      string    `json:"messageType"`
	RoomID           string    `json:"roomID"`
	SenderAvatar     string    `json:"senderAvatar"`
	SenderID         uint64    `json:"senderID"`
	SenderName       string    `json:"senderName"`
	SenderNickname   string    `json:"senderNickname"`
	SendTime         time.Time `json:"sendTime"`
	IsRead           uint64    `json:"isRead"`
	ReceiverID       uint64    `json:"-"`
}

type OfflineMessage struct {
	ID          int       `db:"F_ID"`
	MsgID       uint64    `db:"F_MsgID"`
	SendID      uint64    `json:"sendID" form:"sendID" db:"F_SendID" binding:"required"`
	RecipientID uint64    `json:"recipientID" form:"recipientID" db:"F_RecipientID" binding:"required"`
	Content     string    `json:"content" form:"content" db:"F_Content" binding:"required"`
	SendTime    time.Time `db:"F_SendTime"`
}

type SendMsg struct {
	MsgID    uint64
	SendID   uint64
	SendTime time.Time
	Content  string
}

type OfflineMsg struct {
	SendID      uint64 `json:"sendID" form:"sendID" binding:"required"`
	RecipientID uint64 `json:"recipientID" form:"recipientID" binding:"required"`
}
