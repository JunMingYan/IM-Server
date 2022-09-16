package websocket

const (
	MsgFriend = "FRIEND"
	MsgGroup  = "GROUP"

	MsgH = 1024 //
	MsgR = 1025 //
	MsgA = 1026 //
	MsgN = 1027 //
	AckR = 1028 //
	AckA = 1029
	AckN = 1030

	GroupR = 1031
	GroupA = 1032
	GroupN = 1033
)

type Message struct {
	//MsgID     uint64               `json:"msgID"`
	//MsgType   int                  `json:"msgType"`
	//Topic     string               `json:"topic"`
	Type int         `json:"type"`
	Data interface{} `json:"data"`
	//Ack       uint64               `json:"ack"`
	//HeartBeat string               `json:"heartBeat"`
	//GroupMsg  models.GroupMessage  `json:"groupMsg"`
	//FriendMsg models.SingleMessage `json:"friendMsg"`
}
