package models

type Province struct {
	UserID uint64 `gorm:"column:UserID;type:string;not null;size:32" json:"userId" form:"userId"`
	Name   string `gorm:"column:Name;type:string;not null;size:20" json:"name" form:"name"`
	Code   string `gorm:"column:Code;type:string;not null;size:20" json:"code" form:"code"`
}

func (Province) TableName() string {
	return "province"
}

type City struct {
	UserID uint64 `gorm:"column:UserID;type:string;not null;size:32" json:"userId" form:"userId"`
	Name   string `gorm:"column:Name;type:string;not null;size:20" json:"name" form:"name"`
	Code   string `gorm:"column:Code;type:string;not null;size:20" json:"code" form:"code"`
}

func (City) TableName() string {
	return "city"
}

type Town struct {
	UserID uint64 `gorm:"column:UserID;type:string;not null;size:32" json:"userId" form:"userId"`
	Name   string `gorm:"column:Name;type:string;not null;size:20" json:"name" form:"name"`
	Code   string `gorm:"column:Code;type:string;not null;size:20" json:"code" form:"code"`
}

func (Town) TableName() string {
	return "town"
}
