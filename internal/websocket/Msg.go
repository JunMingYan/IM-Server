package websocket

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Msg struct {
	ID       int
	TopicLen int
	Topic    string
	MsgType  int //1.consumer 2.producer 3.consumer_ack 4.error
	Len      int
	Payload  []byte
}

func BytesToMsg(reader io.Reader) *Msg {
	msg := Msg{}
	var buf [128]byte
	n, err := reader.Read(buf[:])
	if err != nil {
		fmt.Println("read failed,err:", err)
	}
	fmt.Println("read bytes:", n)
	// ID
	buff := bytes.NewBuffer(buf[0:8])
	binary.Read(buff, binary.LittleEndian, &msg.ID)
	fmt.Println(buff)
	// topicLen
	buff = bytes.NewBuffer(buf[8:16])
	binary.Read(buff, binary.LittleEndian, &msg.TopicLen)
	// topic
	msgLastIndex := 16 + msg.TopicLen
	msg.Topic = string(buf[16:msgLastIndex])
	// msgType
	buff = bytes.NewBuffer(buf[msgLastIndex : msgLastIndex+8])
	binary.Read(buff, binary.LittleEndian, &msg.MsgType)

	buff = bytes.NewBuffer(buf[msgLastIndex : msgLastIndex+16])
	binary.Read(buff, binary.LittleEndian, &msg.Len)

	fmt.Printf("broker:-->%v\n", msg)
	if msg.Len <= 0 {
		return &msg
	}

	msg.Payload = buf[msgLastIndex+16:]

	fmt.Printf("broker:-->%v\n", msg)
	return &msg
}

func MsgToBytes(msg *Msg) []byte {
	fmt.Printf("produce->%v\n", msg)

	msg.TopicLen = len([]byte(msg.Topic))
	msg.Len = len(msg.Payload)

	var data []byte
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, msg.ID)
	data = append(data, buf.Bytes()...)

	buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, msg.TopicLen)
	data = append(data, buf.Bytes()...)

	data = append(data, []byte(msg.Topic)...)

	buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, msg.MsgType)
	data = append(data, buf.Bytes()...)

	buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, msg.Len)
	data = append(data, buf.Bytes()...)
	data = append(data, msg.Payload...)
	fmt.Printf("produce->%v\n", buf)

	return data
}
