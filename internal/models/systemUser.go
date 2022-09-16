package models

type SystemUser struct {
	ID       uint64 `gorm:"column:UserID;type:string;not null;size:32" json:"userID" from:"userID"`
	Code     string `gorm:"column:Code;type:string;not null;size:20" json:"code" from:"code"`
	NickName string `gorm:"column:NickName;type:string;not null;size:20" json:"nickName" from:"nickName"`
	Status   int    `gorm:"column:Status;type:int;not null" json:"status" from:"status"`
}

func (SystemUser) TableName() string {
	return "systemUser"
}
