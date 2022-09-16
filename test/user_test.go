package test

//func TestUser(t *testing.T) {
//	// 设置客户端连接配置
//	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//
//	// 连接到MongoDB
//	client, err := mongo.Connect(context.TODO(), clientOptions)
//	if err != nil {
//		logrus.Fatal(err)
//	}
//	DB = client.Database("chat")
//
//	// 检查连接
//	err = client.Ping(context.TODO(), nil)
//	if err != nil {
//		logrus.Fatal(err)
//	}
//
//	user := &models.User{
//		UserName: "JULIA",
//		UserID:   123,
//	}
//	AddUser(user)
//
//	list, err := UserList(bson.M{"userName": "JULIA"})
//	logrus.Infof("list: %v\n", len(list))
//
//	DeleteUser(bson.D{{"username", "JULIA"}})
//}
