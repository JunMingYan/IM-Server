package mysql

import (
	"gorm.io/gorm"
	models "server/internal/models"
)

const (
	secter     string = "JULIA"
	MODIFY_PWD int    = 1
	INVALID_ID int    = -1
)

func SaveUser(user *models.User) error {
	err := DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckExistByUserName(userName string) (models.User, error) {
	var result models.User
	err := DB.Where("UserName = ?", userName).First(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return models.User{}, err
	}
	return result, nil
}

func FindSimpleUser(userName string) (models.SimpleUser, error) {
	var result models.SimpleUser
	err := DB.Table("user").
		Select("UserID,Photo,Signature,NickName,OnlineTime,LastLoginTime,UserName").
		Where("UserName = ?", userName).Scan(&result).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return models.SimpleUser{}, err
	}
	return result, nil
}

func UsrInfo(userID uint64) (models.User, error) {
	var result models.User
	err := DB.Table("user").
		Select("*").
		Where("UserID = ?", userID).Scan(&result).Error
	if err != nil { //  && err != gorm.ErrRecordNotFound
		return models.User{}, err
	}
	return result, nil
}

func SearchUser(user models.SearchUser) ([]models.User, error) {
	var results []models.User
	if user.PageIndex-1 > 0 {
		user.PageIndex = user.PageIndex - 1
	}
	//SELECT * FROM `user` WHERE
	//UserName LIKE CONCAT('%', "A", '%')
	//OR NickName LIKE CONCAT('%', "A", '%')
	//OR `Code` LIKE CONCAT('%', "A", '%')
	//ORDER BY ID DESC LIMIT 0,10;
	err := DB.Table("user").
		Select("*").
		Where("UserName LIKE CONCAT('%', ?, '%') "+
			"OR NickName LIKE CONCAT('%', ?, '%') "+
			"OR `Code` LIKE CONCAT('%', ?, '%')"+
			"ORDER BY ID DESC LIMIT ?,?", user.Key, user.Key, user.Key, user.PageIndex*user.PageSize, user.PageSize).
		Scan(&results).Error
	if err != nil { //  && err != gorm.ErrRecordNotFound
		return nil, err
	}
	return results, nil
}

func UpdatePwd(userID uint64, salt string) error {
	err := DB.Model(models.User{}).
		Where("UserID = ?", userID).
		Update("salt", salt).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserInfo(user models.UpdateUserInfo) error {
	err := DB.Model(models.User{}).
		Where("UserID = ?", user.UserID).
		Update(user.Field, user.Value).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserConfigure(user models.UpdateUserConfigure) error {
	err := DB.Model(models.User{}).
		Where("UserID = ?", user.UserID).
		Updates(models.User{
			BgColor:        user.BgColor,
			BgImg:          user.BgImg,
			Blur:           user.Blur,
			Color:          user.Color,
			CustomBgImgUrl: user.CustomBgImgUrl,
			NotifySound:    user.NotifySound,
			Opacity:        user.Opacity,
		}).Error
	if err != nil {
		return err
	}

	return nil
}
