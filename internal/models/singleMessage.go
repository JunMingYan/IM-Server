package models

import "time"

type SingleMessage struct {
	ID             int       `gorm:"column:ID;primaryKey;" json:"-" form:"-"`
	MsgID          uint64    `gorm:"column:MsgID;type:string;not null;size:20"json:"msgID"`
	RoomID         uint64    `gorm:"column:RoomID;type:string;not null;size:40" json:"roomID" form:"roomID"`
	ReceiverID     uint64    `gorm:"column:ReceiverID;type:string;not null;size:20" json:"receiverID" form:"receiverID"`             // 房间
	SenderID       uint64    `gorm:"column:SenderID;type:string;not null;size:20" json:"senderID" form:"senderID"`                   // 发送者Id
	SenderName     string    `gorm:"column:SenderName;type:string;not null;size:32" json:"senderName" form:"senderName"`             // 发送者登录名
	SenderNickname string    `gorm:"column:SenderNickname;type:string;not null;size:32" json:"senderNickname" form:"senderNickname"` // 发送者昵称
	SenderAvatar   string    `gorm:"column:SenderAvatar;type:string;not null;size:32" json:"senderAvatar" form:"senderAvatar"`       // 发送者头像
	SendTime       time.Time `gorm:"column:SendTime;type:time;not null;size:32" json:"sendTime" form:"sendTime"`                     // 消息发送时间
	FileRawName    string    `gorm:"column:FileRawName;type:string;not null;size:10" json:"fileRawName" form:"fileRawName"`          //文件的原始名字
	Message        string    `gorm:"column:Message;type:string;not null;size:256" json:"message" form:"message"`                     // 消息内容
	MessageType    string    `gorm:"column:MessageType;type:string;not null;size:10" json:"messageType" form:"messageType"`          // 消息的类型：emoji/text/img/file/sys          // 消息的类型：emoji/text/img/file/sys/whiteboard/video/audio
	//IsReadUser       string    `gorm:"column:IsReadUser;type:string;not null;size:40" json:"-" form:"-"`                               // 值为用户的ID，判断已经读取的用户，在发送消息的时候默认发送发已经读取，在单独会话中Array值只有两个
	//IsRead           []uint64  `gorm:"-" json:"-"`
	ConversationType string `gorm:"-" json:"conversationType"`
	IsRead           int    `gorm:"column:IsRead;type:string;not null;size:1" json:"read"`
}

func (SingleMessage) TableName() string {
	return "singleMessage"
}

type RecentSingleMessages struct {
	ReceiverID uint64 `json:"-" form:"-"` // 房间
	SenderID   uint64 `json:"-" form:"-"`
	RoomID     uint64 `json:"roomID" form:"roomID"`
	Key        uint64 `json:"-" `
	PageIndex  int    `json:"pageIndex" form:"pageIndex"`
	PageSize   int    `json:"pageSize" form:"pageSize"`
}

type FilterMessage struct {
	Message    string      `json:"message"`
	RoomID     interface{} `json:"roomID"` //uint64
	SendTime   time.Time   `json:"sendTime"`
	senderID   uint64      `json:"senderID"`
	SenderName string      `json:"senderName"`
	Type       string      `json:"type"`
}

type LastFriendMsg struct {
	RoomID string `json:"roomID" form:"roomID"`
}

type IsRead struct {
	UserID uint64 `json:"userID"`
	RoomID uint64 `json:"roomID"`
}
