package models

import "time"

type FeedBack struct {
	ID              int64     `gorm:"column:ID;primaryKey;" json:"ID"`                                             //主键
	UserID          int64     `gorm:"column:UserID;type:string;not null;size:20" json:"userID"`                    //反馈人id
	UserName        string    `gorm:"column:UserName;type:string;not null;size:20" json:"userName"`                //反馈人账号名
	FeedBackContent string    `gorm:"column:FeedBackContent;type:string;not null;size:256" json:"feedBackContent"` //反馈内容
	CreateTime      time.Time `gorm:"column:CreateTime;type:time;not null;size:20" json:"createTime"`              //反馈时间
}

func (FeedBack) TableName() string {
	return "feedBack"
}
