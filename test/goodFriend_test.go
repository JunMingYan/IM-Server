package test

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongodb2 "server/internal/mongodb"
	"testing"
)

func TestGoodFriend(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	mongodb2.DB = client.Database("chat")

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}

	//gf := models.GoodFriend{
	//	UserM:      "8353560479662081",
	//	UserY:      "8216624960110593",
	//	CreateDate: time.Now(),
	//}
	//
	//err := AddGoodFriend(gf)
	//if err != nil {
	//	t.Fatal(err)
	//}

	var id = "8216624960110593"
	var ids = []string{"8322886209110017", "8353560479662081"}

	results, err := mongodb2.RecentConversation(id, ids)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(results))
}
