package mongodb

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/internal/models"
)

func AddSystemUser(user models.SystemUser) error {
	collection := getCollection(SystemUser)
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Inserted a single document: %v\n", insertResult.InsertedID)
	return err
}

func SystemUserList(M bson.M) ([]*models.SystemUser, error) {
	collection := getCollection(SystemUser)

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

	var result []*models.SystemUser
	if err = cursor.All(context.TODO(), &result); err != nil {
		logrus.Fatal(err)
	}
	return result, err
}
