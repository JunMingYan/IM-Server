package main

import (
	"bytes"
	"fmt"
	"net"
	"server/internal/websocket"
)

// 消费者
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("connect failed,err:", err)
	}
	defer conn.Close()

	msg := websocket.Msg{Topic: "topic-test", MsgType: 1}

	n, err := conn.Write(websocket.MsgToBytes(&msg))
	if err != nil {
		fmt.Println("write failed,err:", err)
	}
	fmt.Printf("consumer->n->%v\n", n)

	var res [128]byte
	conn.Read(res[:])
	buf := bytes.NewBuffer(res[:])
	receMsg := websocket.BytesToMsg(buf)
	fmt.Printf("consumer-%v", receMsg)
	//ack
	conn, _ = net.Dial("tcp", "127.0.0.1:12345")
	l, e := conn.Write(websocket.MsgToBytes(&websocket.Msg{ID: receMsg.ID, Topic: receMsg.Topic, MsgType: 3}))
	if e != nil {
		fmt.Println("Write failed,err:", e)
	}
	fmt.Printf("consumer->l->%v\n", l)
}
