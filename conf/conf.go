package conf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

type Config struct {
	Server  `ini:"websocket"`
	MySql   `ini:"mysql"`
	Redis   `ini:"redis"`
	MongoDB `ini:"mongodb"`
}

type Server struct {
	Mode    string `ini:"mode"`
	Address string `ini:"address"`
	Port    string `ini:"port"`
}

type MySql struct {
	Address string `ini:"address"`
	Port    string `ini:"port"`
	User    string `ini:"user"`
	Pwd     string `ini:"pwd"`
	DBName  string `ini:"dbName"`
}

type Redis struct {
	Address string `ini:"address"`
	Pwd     string `ini:"pwd"`
	DB      int    `ini:"db"`
}

type MongoDB struct {
	Address  string `ini:"address"`
	Port     string `ini:"port"`
	DataBase string `ini:"dataBase"`
}

// var (
// 	App = new(Config)
// )

func Init() (App Config, err error) {
	cfg, err := ini.Load("./conf/conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//
	App.Server.Mode = cfg.Section("websocket").Key("mode").String()
	App.Server.Address = cfg.Section("websocket").Key("address").String()
	App.Server.Port = cfg.Section("websocket").Key("port").String()
	//
	App.MySql.Address = cfg.Section("mysql").Key("address").String()
	App.MySql.Port = cfg.Section("mysql").Key("port").String()
	App.MySql.User = cfg.Section("mysql").Key("user").String()
	App.MySql.Pwd = cfg.Section("mysql").Key("pwd").String()
	App.MySql.DBName = cfg.Section("mysql").Key("dbName").String()
	//
	App.Redis.Address = cfg.Section("redis").Key("address").String()
	App.Redis.Pwd = cfg.Section("redis").Key("pwd").String()
	db, _ := strconv.Atoi(cfg.Section("redis").Key("db").String())
	App.Redis.DB = db
	//
	App.MongoDB.Address = cfg.Section("mongodb").Key("address").String()
	App.MongoDB.Port = cfg.Section("mongodb").Key("port").String()
	App.MongoDB.DataBase = cfg.Section("mongodb").Key("dataBase").String()
	//
	logrus.Infoln(App)
	//
	return App, err
}
