package models

import "time"

type Group struct { //
	ID           int       `gorm:"column:ID;primaryKey;" json:"-"`                               // 主键
	GroupID      int64     `gorm:"column:GroupID;unique;type:string;not null;size:32" json:"id"` // 群ID,自己生成
	Title        string    `gorm:"column:Title;type:string;not null;size:20" json:"title"`       // 群名称
	Desc         string    `gorm:"column:Desc;type:string;not null;size:20" json:"desc"`         //群描述
	Img          string    `gorm:"column:Img;type:string;not null;size:20" json:"img"`           //= "/img/zwsj5.png"//群图片
	Code         string    `gorm:"column:Code;type:string;not null;size:20" json:"code"`         //群号，唯一标识
	UserNum      int       `gorm:"column:UserNum;type:string;not null" json:"userNum"`           // 群成员数量，避免某些情况需要多次联表查找，如搜索；所以每次加入一人，数量加一
	CreateDate   time.Time `gorm:"column:CreateDate;type:time;not null" json:"createDate"`
	HolderName   string    `gorm:"column:HolderName;type:string;not null;size:20" json:"holderName"`     // 群主账号，在user实体中对应name字段
	HolderUserID int64     `gorm:"column:HolderUserID;type:string;not null;size:32" json:"holderUserID"` //群人员的id，作为关联查询
}

func (Group) TableName() string {
	return "group"
}

type GroupMessage struct {
	ID               uint64    `gorm:"column:ID;primaryKey;" json:"id"`
	MsgID            uint64    `gorm:"column:MsgID;type:string;not null;size:20" json:"msgID"`
	RoomID           uint64    `gorm:"column:RoomID;index:idx_roomID;type:string;not null;size:20" json:"roomID"`
	SenderID         uint64    `gorm:"column:SenderID;type:string;not null;size:20" json:"senderID"`
	SenderName       string    `gorm:"column:SenderName;type:string;not null;size:20" json:"senderName"`
	SenderNickName   string    `gorm:"column:SenderNickName;type:string;not null;size:20" json:"senderNickName"`
	SenderAvatar     string    `gorm:"column:SenderAvatar;type:string;not null;size:20" json:"senderAvatar"`
	SendTime         time.Time `gorm:"column:SendTime;type:time;not null;size:20" json:"sendTime"`
	FileRawName      string    `gorm:"column:FileRawName;type:string;not null;size:20" json:"fileRawName"`
	Message          string    `gorm:"column:Message;type:string;not null;size:20" json:"message"`
	MessageType      string    `gorm:"column:MessageType;type:string;not null;size:20" json:"messageType"` // 消息的类型：emoji/text/img/file/sys
	IsReadUser       string    `gorm:"column:IsReadUser;type:string;not null;size:20" json:"-"`
	IsRead           []uint64  `gorm:"-" json:"isReadUser"`
	ConversationType string    `gorm:"-" json:"conversationType"`
}

func (GroupMessage) TableName() string {
	return "groupMessage"
}

type GroupUser struct {
	ID       int64  `gorm:"column:ID;primaryKey" json:"id"`
	GroupID  int64  `gorm:"column:GroupID;type:string;not null;size:32" json:"groupID"`
	UserID   int64  `gorm:"column:UserID;type:string;not null;size:32" json:"userID"`     //成员id
	UserName string `gorm:"column:UserName;type:string;not null;size:20" json:"userName"` //成员账号名
	//Manager  int       `gorm:"column:Manage;type:int;not null" json:"Manager"`               // 是否是管理员，默认0，不是，1是（可以设置一下这个需求）
	//Holder   int       `gorm:"column:Holder;type:int;not null" json:"Holder"`                // 是否是群主，默认0，不是，1是
	Role int       `gorm:"column:Role;type:int;not null" json:"role"`            // 0为群成员,1为管理员,2为群主
	Card string    `gorm:"column:Card;type:string;not null;size:32" json:"card"` // 群名片
	Time time.Time `gorm:"column:Time;type:time;not null" json:"time"`           // 设置默认时间
}

func (GroupUser) TableName() string {
	return "groupUser"
}

type MyGroup struct {
	ID        int64      `json:"-"`
	GroupID   int64      `json:"id"` // 群ID
	UserID    int64      `json:"userID"`
	UserName  string     `json:"userName"`
	Role      int        `json:"role"`
	Card      string     `json:"card"`
	Time      time.Time  `json:"time"`
	GroupInfo Group      `json:"groupInfo"  gorm:"-"`
	UserInfo  SimpleUser `json:"userInfo" gorm:"-"`
	//Manager   int        `json:"manager"`
	//Holder    int        `json:"holder"`
}

type MyGroupForm struct {
	UserName string `json:"userName" form:"userName"`
}

type GroupInfoForm struct {
	GroupID uint64 `json:"groupID" form:"groupID"`
}

type RecentGroupForm struct {
	GroupIDList []int64 `json:"groupIDs"`
	UserID      int64   `json:"userID"`
}

type RecentGroup struct {
	ID        string    `json:"id"`
	UserID    int64     `json:"userId"`
	UserName  string    `json:"username"`
	Manager   int       `json:"manager"`
	Holder    int       `json:"holder"`
	Card      string    `json:"card"`
	Time      time.Time `json:"time"`
	GroupList []Group   `json:"groupList" gorm:"-"`
}

type RecentGroupMessage struct {
	RoomID    uint64 `json:"roomID"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}
