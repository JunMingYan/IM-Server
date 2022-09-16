package models

import "time"

type BrowserSetting struct {
	UserID    uint64    `gorm:"column:UserID;unique;type:string;not null;size:32"`
	Browser   string    `gorm:"column:Browser;type:string;not null;size:20" json:"browser"`
	Country   string    `gorm:"column:Country;type:string;not null;size:20" json:"country"`
	IP        string    `gorm:"column:IP;type:string;not null;size:20" json:"ip"`
	OS        string    `gorm:"column:OS;type:string;not null;size:20" json:"os"`
	LoginTime time.Time `gorm:"column:LoginTime;type:time;not null" json:"LoginTime"`
}

func (BrowserSetting) TableName() string {
	return "browserSetting"
}

type Setting struct {
	Browser string `json:"browser"`
	Country string `json:"country"`
	IP      string `json:"IP"`
	OS      string `json:"OS"`
}
