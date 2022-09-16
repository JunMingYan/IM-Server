package mysql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"server/conf"
	models "server/internal/models"
)

var (
	DB *gorm.DB
)

func Init(app *conf.Config) (err error) {
	strDNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		app.MySql.User, app.MySql.Pwd, app.MySql.Address, app.MySql.Port, app.MySql.DBName)
	logrus.Infoln(strDNS)
	DB, err = gorm.Open(mysql.Open(strDNS), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	DB.AutoMigrate(
		&models.User{},
		&models.SystemUser{},
		&models.Conversation{},
		&models.GoodFriend{},
		//&models.BrowserSetting{},
		//&models.Province{},
		//&models.City{},
		//&models.Town{},
		&models.FriendGroup{},
		&models.FriendNote{},
		&models.Group{},
		&models.GroupUser{},
		&models.ValidateMessage{},
		&models.FriendGroup{},
		&models.FriendNote{},
		&models.SingleMessage{},
		&models.GroupMessage{},
		&models.FeedBack{},
		&models.FriendGroupName{},
	)
	//
	return
}
