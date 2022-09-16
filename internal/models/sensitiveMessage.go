package models

type SensitiveMessage struct {
	ID         int64  //发送者编号
	RoomId     string //房间号
	SenderId   string // 发送者Id
	SenderName string // 发送者登录名
	Time       string // 消息发送时间
	Message    string // 消息内容
	Type       string // 0/1/2, 好友/群聊/验证消息
}
