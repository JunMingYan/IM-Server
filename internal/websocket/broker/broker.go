package main

import (
	"bufio"
	"container/list"
	"fmt"
	"net"
	"os"
	"server/internal/websocket"
	"sync"
	"time"
)

type Queue struct {
	len  int
	data list.List
}

var lock sync.Mutex

func (queue *Queue) offer(msg *websocket.Msg) {
	queue.data.PushBack(msg)
	queue.len = queue.data.Len()
}

func (queue *Queue) poll() *websocket.Msg {
	if queue.len == 0 {
		return nil
	}
	msg := queue.data.Front()
	return msg.Value.(*websocket.Msg)
}

func (queue *Queue) delete(id int) {
	lock.Lock()
	for msg := queue.data.Front(); msg != nil; msg = msg.Next() {
		if msg.Value.(websocket.Msg).ID == id {
			queue.data.Remove(msg)
			queue.len = queue.data.Len()
		}
	}
	lock.Unlock()
}

var topics = sync.Map{}

func handleErr(conn net.Conn) {
	if err := recover(); err != nil {
		println(err.(string))
		conn.Write(websocket.MsgToBytes(&websocket.Msg{MsgType: 4}))
	}
}

func Process(conn net.Conn) {
	defer handleErr(conn)
	reader := bufio.NewReader(conn)
	msg := websocket.BytesToMsg(reader)
	queue, ok := topics.Load(msg.Topic)
	var res *websocket.Msg
	if msg.MsgType == 1 {
		// consumer
		if queue == nil || queue.(*Queue).len == 0 {
			return
		}
		msg = queue.(*Queue).poll()
		msg.MsgType = 1
		res = msg
	} else if msg.MsgType == 2 {
		// producer
		if !ok {
			queue = &Queue{}
			queue.(*Queue).data.Init()
			topics.Store(msg.Topic, queue)
		}
	} else if msg.MsgType == 3 {
		// consumer ack
		if queue == nil {
			return
		}
		queue.(*Queue).delete(msg.ID)
	}
	conn.Write(websocket.MsgToBytes(res))
}

func Save() {
	ticker := time.NewTicker(60)
	for {
		select {
		case <-ticker.C:
			topics.Range(func(key, value interface{}) bool {
				if value == nil {
					return false
				}
				file, _ := os.Open(key.(string))
				if file == nil {
					file, _ = os.Create(key.(string))
				}
				for msg := value.(*Queue).data.Front(); msg != nil; msg = msg.Next() {
					file.Write(websocket.MsgToBytes(msg.Value.(*websocket.Msg)))
				}
				file.Close()
				return false
			})
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}

	go Save()
	for {
		conn, err := listen.Accept()
		fmt.Printf("有新连接进来->%v\n", conn)
		if err != nil {
			fmt.Println("accept failed err:", err)
			continue
		}
		go Process(conn)
	}
}
