package websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"server/internal/models"
)

type Client struct {
	Conn    *websocket.Conn
	User    models.UserWS
	Send    chan []byte
	MsgType chan int
}

func (c *Client) Read() {
	defer func() {
		MyServer.Ungister <- c
		c.Conn.Close()
		if err := recover(); err != nil {
			logrus.Info(err)
			return
		}
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			//panic(err)
			c.Conn.Close()
			return
		}

		msg, err := UnmarshalMsg(message)
		if err != nil {
			panic(err)
			continue
		}

		if msg.Type == MsgH {
			MyServer.HeartBeat <- c
		}

		if msg.Type == MsgR || msg.Type == GroupR {
			MyServer.Broadcast <- msg
		}

		if msg.Type == AckR {
			MyServer.AckResp <- msg
		}

	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()
	for message := range c.Send {
		MyServer.mutex.Lock()
		c.Conn.WriteMessage(websocket.TextMessage, message)
		MyServer.mutex.Unlock()
	}
}

func UnmarshalMsg(buffer []byte) (Message, error) {
	data := Message{}
	err := json.Unmarshal(buffer, &data)
	if err != nil {
		return Message{}, err
	}
	return data, nil
}
