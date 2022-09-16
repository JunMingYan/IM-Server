package models

type Conversation struct {
	Name  string `bson:"name" gorm:"column:Name;type:string;not null;size:20"`
	Photo string `bson:"photo" gorm:"column:Photo;type:string;not null;size:20"`
	ID    uint64 `bson:"ID" gorm:"column:ID;type:string;not null;size:32"`
	Type  string `bson:"type" gorm:"column:Type;type:string;not null;size:20"` // 会话类型 group/ frend
}

func (Conversation) TableName() string {
	return "conversation"
}

type ConversationForm struct {
	UserID          uint64   `json:"userID" form:"userID" binding:"required"`
	RecentFriendIds []uint64 `json:"friendList" form:"friendList" binding:"required"`
}
