package mongodb

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/internal/models"
)

func AddConversation(c models.Conversation) error {
	collection := getCollection(Conversation)
	insertResult, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Inserted a single document: %v\n", insertResult.InsertedID)
	return err
}

func ConversationList(M bson.M) ([]*models.Conversation, error) {
	collection := getCollection(Conversation)

	findOptions := options.Find()
	findOptions.SetLimit(2)

	cursor, err := collection.Find(context.TODO(), M, findOptions)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			logrus.Fatal(err)
		}
	}()

	//var results []*models.User
	//for cursor.Next(context.TODO()) {
	//	var elem models.User
	//	err := cursor.Decode(&elem)
	//	if err != nil {
	//		logrus.Fatal(err)
	//	}
	//	results = append(results, &elem)
	//}

	var result []*models.Conversation
	if err = cursor.All(context.TODO(), &result); err != nil {
		logrus.Fatal(err)
	}
	return result, err
}

func test() {
	var ids []int
	var id int
	collection := getCollection(GoodFriend)

	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", []bson.E{
				{"userM", bson.D{{"$in", bson.A{ids}}}},
				{"userY", id},
			},
			},
		},
		bson.D{
			{"$match", []bson.E{
				{"userY", bson.D{{"$in", bson.A{ids}}}},
				{"userM", id},
			},
			},
		},
		bson.D{
			{
				"$lookup", bson.M{
					"from":         User,
					"localField":   "userM",
					"foreignField": "userID",
					"as":           "UserList1",
				},
			},
		},
		bson.D{
			{
				"$lookup", bson.M{
					"from":         User,
					"localField":   "userY",
					"foreignField": "userID",
					"as":           "UserList2",
				},
			},
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		logrus.Fatal(err)
	}

	var results []models.MyFriend
	if err = cursor.All(context.TODO(), &results); err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(len(results))
}
