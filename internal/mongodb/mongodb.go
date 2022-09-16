package mongodb

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server/conf"
)

var DB *mongo.Database

var client *mongo.Client

const (
	User         = "user"
	SystemUser   = "systemUser"
	Conversation = "conversation"
	GoodFriend   = "goodFriend"
)

func Init(app *conf.Config) (err error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	DB = client.Database(app.MongoDB.DataBase)

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}
	return
}

func Close() (err error) {
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	logrus.Fatal("Connection to MongoDB closed.")
	return
}

func getCollection(collection string) *mongo.Collection {
	return DB.Collection(collection)
}
