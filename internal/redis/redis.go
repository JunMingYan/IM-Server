package redis

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"server/conf"
)

var db *redis.Client

func Init(app *conf.Config) error {
	db = redis.NewClient(&redis.Options{
		Addr:     app.Redis.Address,
		Password: app.Redis.Pwd,
		DB:       app.Redis.DB,
	})
	//
	pong, err := db.Ping().Result()
	if err != nil {
		logrus.Errorf("ping redis 出错,err:%s\n", err)
		return err
	}
	logrus.Infof("ping result: %s\n", pong)
	return nil
}
