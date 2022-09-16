package mongodb

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/internal/models"
)

func RecentConversation(id string, ids []string) ([]models.SingleRecentConversation, error) {
	collection := getCollection(GoodFriend)

	pipeline := mongo.Pipeline{
		bson.D{
			{
				"$match", []bson.E{{"userM", id}, {"userY", bson.D{{"$in", ids}}}},
			},
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	var myFriends []models.GoodFriend
	if err = cursor.All(context.TODO(), &myFriends); err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	var userM, userY models.SimpleUser
	var conversationList []models.SingleRecentConversation
	for _, v := range myFriends {
		userM, err = FindSimpleUser(bson.D{{"userID", v.UserM}})
		if err != nil {
			logrus.Fatal(err)
			return nil, err
		}

		userY, err = FindSimpleUser(bson.D{{"userID", v.UserY}})
		if err != nil {
			logrus.Fatal(err)
			return nil, err
		}
		conversation := models.SingleRecentConversation{
			ID:         v.UserM,
			UserM:      userM,
			UserY:      userY,
			CreateDate: v.CreateDate,
		}

		conversationList = append(conversationList, conversation)
	}

	return conversationList, err
}

func AddGoodFriend(gf models.GoodFriend) error {
	collection := getCollection(GoodFriend)
	insertResult, err := collection.InsertOne(context.TODO(), gf)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Inserted a single document: %v\n", insertResult.InsertedID)
	return err
}
