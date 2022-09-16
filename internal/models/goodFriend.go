package models

import "time"

type GoodFriend struct {
	//ID         int64     `bson:"ID"`
	UserM      uint64    `bson:"userM" gorm:"column:UserM;type:string;not null;size:32"`
	UserY      uint64    `bson:"userY" gorm:"column:UserY;type:string;not null;size:32"`
	CreateDate time.Time `bson:"createDate" gorm:"column:CreateDate;type:time;not null"` // 加好友时间
}

func (GoodFriend) TableName() string {
	return "goodFriend"
}

type MyFriendForm struct {
	UserID uint64 `json:"userID" form:"userID" binding:"required"`
}

type MyFriend struct {
	CreateDate time.Time `json:"createDate" form:"createDate" binding:"required"`
	NickName   string    `json:"nickname" form:"nickname" binding:"required"`
	Photo      string    `json:"photo" form:"photo" binding:"required"`
	Signature  string    `json:"signature" form:"signature" binding:"required"`
	UserID     uint64    `json:"id" form:"id" binding:"required"`
	Level      int       `json:"level" form:"level" binding:"required"`
	RoomID     uint64    `json:"roomID" form:"roomID" binding:"required"`
	OnlineTime int64
}

type SimpleUser struct {
	UserID        uint64    `json:"userID" form:"userID" binding:"required"`
	Photo         string    `json:"photo" form:"photo" binding:"required"`
	Signature     string    `json:"signature" form:"signature" binding:"required"`
	NickName      string    `json:"nickname" form:"nickname" binding:"required"`
	OnlineTime    int64     `json:"onlineTime" form:"onlineTime" binding:"required"`
	Level         int       `json:"level" form:"level" binding:"required"`
	LastLoginTime time.Time `json:"lastLoginTime" form:"lastLoginTime" binding:"required"`
	UserName      string    `json:"userName" form:"userName" binding:"required"`
	CreateDate    time.Time `json:"createDate" form:"createDate" binding:"required"`
}

type SingleRecentConversation struct {
	UserID     uint64     `json:"userID" form:"userID" binding:"required"`
	CreateDate time.Time  `json:"createDate" form:"createDate" binding:"required"`
	UserM      SimpleUser `json:"userM" form:"userM" binding:"required"`
	UserY      SimpleUser `json:"userY" form:"userY" binding:"required"`
	RoomID     uint64     `json:"roomID" form:"roomID"`
}
