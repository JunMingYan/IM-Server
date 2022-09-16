package models

import (
	"time"
)

type User struct { //1.字段不传输 `json:"-"` 2.字段为空不传  `json:"omitempty"`
	ID        int64  `gorm:"column:ID;primaryKey;"`
	UserID    uint64 `gorm:"column:UserID;unique;type:string;not null;size:32" json:"userID" form:"userID"`
	UserName  string `gorm:"column:UserName;unique;type:string;not null;size:20" json:"userName"`
	Salt      string `gorm:"column:Salt;type:string;not null;size:64" json:"-"`
	Avatar    string `gorm:"column:Avatar;type:string;not null;size:20" json:"avatar"`
	Code      string `gorm:"column:Code;type:string;not null;size:20" json:"code"`
	Photo     string `gorm:"column:Photo;type:string;not null;size:20" json:"photo"`
	Signature string `gorm:"column:Signature;type:string;not null;size:20" json:"signature"`
	NickName  string `gorm:"column:NickName;type:string;not null;size:20" json:"nickName"`
	Email     string `gorm:"column:Email;type:string;not null;size:20" json:"email"` // TODO unique
	//Province       Province                 `gorm:"-" json:"province"`
	//City           City                     `gorm:"-" json:"city"`
	//Town           Town                     `gorm:"-" json:"town"`
	Sex            int       `gorm:"column:Sex;type:int;not null" json:"sex"`
	Opacity        float32   `gorm:"column:Opacity;type:float;not null" json:"opacity"`
	Blur           int       `gorm:"column:Blur;type:int;not null" json:"blur"`
	BgImg          string    `gorm:"column:BgImg;type:string;not null;size:20" json:"bgImg"`
	CustomBgImgUrl string    `gorm:"column:CustomBgImgUrl;type:string;not null;size:20" json:"customBgImgUrl"`
	NotifySound    string    `gorm:"column:NotifySound;type:string;not null;size:20" json:"notifySound"`
	Color          string    `gorm:"column:Color;type:string;not null;size:20" json:"color"`
	BgColor        string    `gorm:"column:BgColor;type:string;not null;size:20" json:"bgColor"`
	SignUpTime     time.Time `gorm:"column:SignUpTime;type:time;not null" json:"signUpTime"`
	LastLoginTime  time.Time `gorm:"column:LastLoginTime;type:time;not null" json:"lastLoginTime"`
	Status         int       `gorm:"column:Status;type:int;not null" json:"status"`
	Age            int       `gorm:"column:Age;type:int;not null" json:"age"`
	OnlineTime     int64     `gorm:"column:OnlineTime;type:int;not null" json:"onlineTime"`
	//LoginSetting   BrowserSetting           `gorm:"-" json:"loginSetting"`
	AddFriendTime time.Time `gorm:"-" json:"-"`
	//FriendGroups  map[string][]uint64 `json:"friendFenZu"`
	//FriendNotes   map[uint64]string   `json:"friendBeiZhu"`
}

func (User) TableName() string {
	return "user"
}

type FriendGroup struct {
	ID        int64  `gorm:"column:ID;primaryKey;"`
	UserID    uint64 `gorm:"column:UserID;type:string;not null;size:32" json:"userID"`
	FriendID  uint64 `gorm:"column:FriendID;type:string;not null;size:20" json:"friendID"`
	GroupName string `gorm:"column:GroupName;type:string;not null;size:20" json:"groupName"`
	GroupID   int64  `gorm:"column:GroupID;type:int;not null;size:20;" json:"-"`
}

func (FriendGroup) TableName() string {
	return "friendGroup"
}

type FriendGroupName struct {
	ID        int64  `gorm:"column:ID;primaryKey;"`
	UserID    uint64 `gorm:"column:UserID;type:string;not null;size:32"`
	GroupName string `gorm:"column:GroupName;type:string;not null;size:20"`
}

func (FriendGroupName) TableName() string {
	return "friendGroupName"
}

type FriendNote struct {
	UserID   uint64 `gorm:"column:UserID;type:string;not null;size:32"`
	Notes    string `gorm:"column:Notes;type:string;not null;size:20"`
	FriendID uint64 `gorm:"column:FriendID;type:string;not null;size:20"`
	//FriendName string `gorm:"column:FriendName;type:string;not null;size:20""`
}

func (FriendNote) TableName() string {
	return "friendNote"
}

type UpdateFriendNote struct {
	FriendBeiZhuName string `json:"friendBeiZhuName"`
	FriendID         uint64 `json:"friendID"`
	UserID           uint64 `json:"userID"`
}

type UpdateFriendGroup struct {
	FriendGroup string `json:"newFenZuName"`
	FriendID    uint64 `json:"friendID"`
	UserID      uint64 `json:"userID"`
}

type UpdateFriendGroupName struct {
	NewFenZu string `json:"newFenZu"`
	OldFenZu string `json:"oldFenZu"`
	UserID   uint64 `json:"userID"`
}

type AddFriendGroup struct {
	FenZuName string `json:"fenZuName"`
	UserID    uint64 `json:"userID"`
}

type DeleteFriendGroup struct {
	UserID    uint64 `json:"userID"`
	FenZuName string `json:"fenZuName"`
}

type UserRegisterForm struct {
	UserName         string `json:"username" form:"username" binding:"required"`
	Password         string `json:"password" form:"password" binding:"required"`
	Confirm_Password string `json:"rePassword" form:"rePassword" binding:"required"`
	Code             string `json:"cvCode" form:"cvCode" binding:"required"`
	Avatar           string `json:"avatar" form:"avatar" binding:"required"`
}

type UserLoginForm struct {
	UserName string  `json:"username" form:"username" binding:"required"`
	Password string  `json:"password" form:"password" binding:"required"`
	Code     string  `json:"cvCode" form:"cvCode" binding:"required"`
	Avatar   string  `json:"avatar"`
	Setting  Setting `json:"setting"`
}

//type UserInfo struct {
//	UserName string `gorm:"column:UserName;unique"`
//	Avatar   string `gorm:"column:Avatar"`
//}

type UserInfoForm struct {
	UserID uint64 `json:"userID" form:"userID"`
}

type SearchUser struct {
	Key       string `json:"searchContent" form:"searchContent"`
	PageIndex int    `json:"pageIndex" form:"pageIndex"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
	Type      string `json:"type" form:"type"`
}

type UserWS struct { // http://localhost:8899/socket.io/?EIO=3&transport=polling&t=OAt1JZQ
	ID uint64 `json:"ID" form:"ID" binding:"required"` //
	//UserName string `json:"username" form:"username" binding:"required"` //
}

type WS struct {
	RoomID          string    `json:"roomID" form:"roomID"`
	SenderID        int64     `json:"senderID" form:"senderID"`
	SenderName      string    `json:"senderName" form:"senderName"`
	SenderNickName  string    `json:"senderNickName" form:"senderNickName"`
	SenderAvatar    string    `json:"senderAvatar" form:"senderAvatar"`
	ReceiverID      int64     `json:"receiverID" form:"receiverID"`
	Time            time.Time `json:"time" form:"time"`
	AdditionMessage string    `json:"additionMessage" form:"additionMessage"`
	Status          int       `json:"status" form:"status"`
	ValidateType    int       `json:"validateType" form:"validateType"`
}

type UpdateUserPwd struct {
	NewPwd   string
	OldPwd   string
	ReNewPwd string
	UserID   uint64
}

type UpdateUserInfo struct {
	Field  string `json:"field"`
	UserID uint64 `json:"userID"`
	Value  string `json:"value"`
}

type UpdateUserConfigure struct {
	BgColor        string  `json:"bgColor"`
	BgImg          string  `json:"bgImg"`
	Blur           int     `json:"blur"`
	Color          string  `json:"color"`
	CustomBgImgUrl string  `json:"customBgImgUrl"`
	NotifySound    string  `json:"notifySound"`
	Opacity        float32 `json:"opacity"`
	UserID         uint64  `json:"userID"`
}

func GetUser() *User {
	return &User{
		Photo:         "static/face/picture.png",
		Sex:           3,
		Opacity:       0.75,
		Blur:          10,
		BgImg:         "abstract",
		NotifySound:   "default",
		Color:         "#000",
		BgColor:       " #fff",
		Status:        0,
		Age:           18,
		SignUpTime:    time.Now(),
		LastLoginTime: time.Now(),
	}
}
