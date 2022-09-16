package main

import (
	"fmt"
	"net"
	"server/internal/websocket"
)

// 生产者
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("connect failed,err:", err)
	}
	defer conn.Close()

	msg := websocket.Msg{ID: 1102, Topic: "topic-test", MsgType: 2, Payload: []byte("JULIA")}
	n, err := conn.Write(websocket.MsgToBytes(&msg))
	if err != nil {
		fmt.Println("Write failed,err:", n)
	}

	fmt.Printf("produce->%v\n", n)
}
