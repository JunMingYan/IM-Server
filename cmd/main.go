package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"server/conf"
	"server/internal/mysql"
	"server/internal/websocket"
	"server/pkg/sonyflake"
	"server/routers"
	"time"
)

func main() {
	App, err := conf.Init()
	if err != nil {
		logrus.Errorf("初始化App出错,err:%s", err)
	}
	//
	err = mysql.Init(&App)
	if err != nil {
		logrus.Errorf("连接数据库出错,err:%s", err)
	}
	//defer mysql.Close()
	//
	err = sonyflake.Init(1)
	if err != nil {
		logrus.Errorf("初始化雪花算法出错,err:%s", err)
	}
	//
	//err = redis.Init(&App)
	//if err != nil {
	//	logrus.Errorf("初始化redis出错,err:%s", err)
	//}
	//
	//err = mongodb.Init(&App)
	//if err != nil {
	//	logrus.Errorf("初始化mongodb出错,err:%s\n", err)
	//}
	//defer mongodb.Close()
	//
	go websocket.MyServer.Start()

	r := routers.GetRouter()
	r.Run(":8899")
	//
	s := &http.Server{
		Addr:           ":8899",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = s.ListenAndServe()
	if nil != err {
		logrus.Errorf("启动服务错误,err=>%s\n", err)
	}

}
