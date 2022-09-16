package models

type AccountPool struct {
	Code   int64 `bson:"code"`   //= 10000000L; //用户或群聊标识（code字段在数据库中还是会以_id的名字存在）
	Type   int   `bson:"type"`   //1：用户；2：群聊
	Status int   `bson:"status"` //1：已使用；0：未使用
}
