package websocket

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	models "server/internal/models"
	service "server/internal/service"
	"server/pkg/sonyflake"
	"sync"
	"time"
)

type Server struct {
	Clients   map[uint64]*Client
	mutex     *sync.Mutex
	Broadcast chan Message // 广播
	Register  chan *Client // 注册
	Ungister  chan *Client // 注销
	HeartBeat chan *Client // 心跳
	AckResp   chan Message
	packer    IPacker
}

var MyServer = NewServer()

func NewServer() *Server {
	return &Server{
		mutex:     &sync.Mutex{},
		Clients:   make(map[uint64]*Client),
		Broadcast: make(chan Message, 1024),
		Register:  make(chan *Client, 1024),
		Ungister:  make(chan *Client, 1024),
		HeartBeat: make(chan *Client, 1024),
		AckResp:   make(chan Message, 1024),
		packer:    &NormalPacker{ByteOrder: binary.BigEndian},
	}
}

func (s *Server) Start() {
	for {
		select {
		case conn := <-s.Register: // 用户注册
			s.Clients[conn.User.ID] = conn
		case conn := <-s.Ungister: // 用户注销
			if _, ok := s.Clients[conn.User.ID]; ok {
				close(conn.Send)
				delete(s.Clients, conn.User.ID)
			}
		case sender := <-s.HeartBeat: // TODO 用Redis实现显示在线
			err := HeartBeat(sender)
			if err != nil {
				logrus.Info(err)
				continue
			}
		case data := <-s.Broadcast: // 广播: 1.私聊 2.群聊
			if data.Type == MsgR {
				message, err := UnmarshaMsg(data)
				if err != nil {
					logrus.Info(err)
					continue
				}
				err = FriendChat(s, message)
				if err != nil {
					logrus.Info(err)
					continue
				}
			}
			if data.Type == GroupR { // TODO 实现群聊逻辑
				message, err := UnmarshaGroup(data)
				if err != nil {
					logrus.Info(err)
					continue
				}

				fmt.Println(len(message.Message))
				fmt.Println(message.Message)

				GroupChat(s, message)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func HeartBeat(sender *Client) error {
	heartBeat := Message{Type: MsgH, Data: "HeartBeat"}
	byte, err := json.Marshal(heartBeat)
	if err != nil {
		return err
	}
	sender.Send <- byte

	return nil
}

func FriendChat(s *Server, message models.SingleMessage) error {
	// 1.消息响应
	sender, ok := s.Clients[message.SenderID]
	if !ok {
		return errors.New("发送者不存在")
	}
	err := MsgResponse(sender, message)
	if err != nil {
		return err
	}
	// 2.查询接收者是否在线
	receiver, ok := s.Clients[message.ReceiverID]
	if !ok { // 对方不在线
		FriendOffline(sender, message)
	} else { // 对方在线
		FriendOnline(s, sender, receiver, message)
	}

	return nil
}

func FriendOnline(s *Server, sender *Client, receiver *Client, message models.SingleMessage) error {
	// 1.存储消息
	message.IsRead = 1
	msg, err := SaveMsg(message)
	if err != nil {
		return err
	}
	// 2.向接收者发送消息请求包
	notice := Message{Type: MsgN, Data: msg}
	dataByte, err := json.Marshal(notice)
	if err != nil {
		return err
	}
	receiver.Send <- dataByte
	// 3.等待接收者返回的消息响应包
	flag := true
	idleDuration := 30 * time.Second
	idleTimeout := time.NewTimer(idleDuration)
	for {
		select {
		case ack := <-s.AckResp:
			// 4.向接收者发送消息发送包
			flag = false
			response := Message{Type: AckA, Data: ack.Data}
			dataByte, err := json.Marshal(response)
			if err != nil {
				return err
			}
			receiver.Send <- dataByte
			// 5.向发送者发送消息通知包(接收者已收到消息)
			notice := Message{Type: AckN, Data: ack.Data}
			dataByte, err = json.Marshal(notice)
			if err != nil {
				return err
			}
			sender.Send <- dataByte
			break
		case <-idleTimeout.C: // 超时
			if flag {
				notice := Message{Type: MsgN, Data: msg}
				dataByte, err := json.Marshal(notice)
				if err != nil {
					return err
				}
				receiver.Send <- dataByte
			}
			break
		default:
			time.Sleep(1 * time.Second)
		}
	}

}

func FriendOffline(sender *Client, message models.SingleMessage) error {
	// 1.存储消息
	message.IsRead = 0
	msg, err := SaveMsg(message)
	if err != nil {
		return err
	}
	// 2.将消息添加到离线消息队列 TODO 实现消息队列
	// 3.向发送者发送消息响应包(接收者已收到消息)
	return AckResponse(sender, msg)
}

func SaveMsg(message models.SingleMessage) (models.SingleMessage, error) {
	//message.IsRead = append(message.IsRead, message.ReceiverID)
	//str, err := json.Marshal(message.IsRead)
	//if err != nil {
	//	return models.SingleMessage{}, nil
	//}
	//
	//message.IsReadUser = string(str)
	msg, err := service.SaveSingleMessages(message)
	if err != nil {
		return models.SingleMessage{}, nil
	}

	return msg, nil
}

func MsgResponse(sender *Client, msg models.SingleMessage) error { //response
	data := make(map[string]interface{})
	data["msgID"] = msg.MsgID
	data["sendTime"] = msg.SendTime

	response := Message{Type: MsgA, Data: data}
	dataByte, err := json.Marshal(response)
	if err != nil {
		return err
	}

	sender.Send <- dataByte

	return nil
}

func AckResponse(sender *Client, msg models.SingleMessage) error { //response
	data := make(map[string]interface{})
	data["msgID"] = msg.MsgID
	data["sendTime"] = msg.SendTime

	response := Message{Type: AckN, Data: data}
	dataByte, err := json.Marshal(response)
	if err != nil {
		return err
	}

	sender.Send <- dataByte

	return nil
}

func UnmarshaMsg(data Message) (models.SingleMessage, error) {
	// 1.解析数据
	msgByte, err := json.Marshal(data.Data)
	if err != nil {
		return models.SingleMessage{}, err
	}

	msg := models.SingleMessage{}
	err = json.Unmarshal(msgByte, &msg)
	if err != nil {
		return models.SingleMessage{}, err
	}
	// 2.生成消息ID和解析接收者
	msg.MsgID, _ = sonyflake.GetMsgID()
	//list := strings.Split(msg.RoomID, "-")
	//ReceiverID, err := strconv.Atoi(list[1])
	//if err != nil {
	//	return models.SingleMessage{}, err
	//}
	//msg.ReceiverID = uint64(ReceiverID)

	return msg, nil
}

func UnmarshaGroup(data Message) (models.GroupMessage, error) {
	// 1.解析数据
	msgByte, err := json.Marshal(data.Data)
	if err != nil {
		return models.GroupMessage{}, err
	}

	msg := models.GroupMessage{}
	err = json.Unmarshal(msgByte, &msg)
	if err != nil {
		return models.GroupMessage{}, err
	}
	// 2.生成消息ID和解析接收者
	msg.MsgID, _ = sonyflake.GetGroupID()
	if err != nil {
		return models.GroupMessage{}, err
	}

	return msg, nil
}

func GroupChat(s *Server, message models.GroupMessage) error {
	// TODO 实现群聊
	// 1.保存群消息
	err := service.SaveGroupMessage(message)
	if err != nil {
		return err
	}
	// 2.向发送者发送消息响应包
	sender, ok := s.Clients[message.SenderID]
	if !ok {
		return errors.New("发送者不存在")
	}
	response := Message{Type: GroupA, Data: "Sent successfully"}
	byte, err := json.Marshal(response)
	if err != nil {
		return err
	}
	sender.Send <- byte
	// 3.向在线的接收者(除发送者外的群成员)发送消息通知包，不在线的把消息添加离线消息队列
	groupUserList, err := service.GroupUserList(message.RoomID)
	if err != nil {
		return err
	}
	request := Message{Type: GroupN, Data: message}
	byte, err = json.Marshal(request)
	if err != nil {
		return err
	}
	for _, v := range groupUserList {
		user, ok := s.Clients[v.UserID]
		if !ok {
			logrus.Infof("用户:%s不在线,添加一条离线消息", v.UserName) // TODO 实现消息队列
			continue
		}
		if v.UserID != message.SenderID {
			user.Send <- byte
		}
	}
	return nil
}
