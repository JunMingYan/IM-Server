package mongodb

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server/internal/models"
)

func AddUser(user *models.User) error {
	collection := getCollection(User)
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Inserted a single document: %v\n", insertResult.InsertedID)
	return err
}

func FindSimpleUser(filter bson.D) (models.SimpleUser, error) {
	collection := getCollection(User)
	var result models.SimpleUser
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		logrus.Fatal(err)
		return models.SimpleUser{}, err
	}
	return result, err
}

func DeleteUser(filter bson.D) error {
	collection := getCollection(User)
	//
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	logrus.Infof("Deleted %v documents in the trainers collection \n", deleteResult.DeletedCount)
	return err
}

func UserList(m bson.M) ([]*models.User, error) {
	collection := getCollection(User)

	findOptions := options.Find()
	findOptions.SetLimit(2)

	cursor, err := collection.Find(context.TODO(), m, findOptions)
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

	var result []*models.User
	if err = cursor.All(context.TODO(), &result); err != nil {
		logrus.Fatal(err)
	}
	return result, err
}

func CheckUser(filter bson.D) bool {
	collection := getCollection(User)
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Found a single document: %v\n", count)
	return count == 0
}
