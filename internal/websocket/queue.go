package websocket

//import (
//	"bufio"
//	"container/list"
//	"os"
//	"sync"
//	"time"
//)
//
//type Queue struct {
//	len  int
//	data list.List
//}
//
//var lock sync.Mutex
//
//func (queue *Queue) offer(msg Message) {
//	queue.data.PushBack(msg)
//	queue.len = queue.data.Len()
//}
//
//func (queue *Queue) poll() Message {
//	if queue.len == 0 {
//		return Message{}
//	}
//	msg := queue.data.Front()
//	return msg.Value.(Message)
//}
//
//func (queue *Queue) delete(id uint64) {
//	lock.Lock()
//	for msg := queue.data.Front(); msg != nil; msg = msg.Next() {
//		if msg.Value.(Message).MsgID == id {
//			queue.data.Remove(msg)
//			queue.len = queue.data.Len()
//		}
//	}
//	lock.Unlock()
//}
//
//var topics = sync.Map{}
//
//func handleErr() {
//	if err := recover(); err != nil {
//		println(err.(string))
//	}
//}
//
//func Process(msg Message) {
//	defer handleErr()
//	reader := bufio.NewReader(conn)
//	msg := websocket.BytesToMsg(reader)
//	queue, ok := topics.Load(msg.Topic)
//	var res Message
//	if msg.MsgType == 1 {
//		// consumer
//		if queue == nil || queue.(*Queue).len == 0 {
//			return
//		}
//		msg = queue.(*Queue).poll()
//		msg.MsgType = 1
//		res = msg
//	} else if msg.MsgType == 2 {
//		// producer
//		if !ok {
//			queue = &Queue{}
//			queue.(*Queue).data.Init()
//			topics.Store(msg.Topic, queue)
//		}
//	} else if msg.MsgType == 3 {
//		// consumer ack
//		if queue == nil {
//			return
//		}
//		queue.(*Queue).delete(msg.ID)
//	}
//	conn.Write(websocket.MsgToBytes(res))
//}
//
//func Save() {
//	ticker := time.NewTicker(60)
//	for {
//		select {
//		case <-ticker.C:
//			topics.Range(func(key, value interface{}) bool {
//				if value == nil {
//					return false
//				}
//				file, _ := os.Open(key.(string))
//				if file == nil {
//					file, _ = os.Create(key.(string))
//				}
//				for msg := value.(*Queue).data.Front(); msg != nil; msg = msg.Next() {
//					file.Write(websocket.MsgToBytes(msg.Value.(websocket.Msg)))
//				}
//				file.Close()
//				return false
//			})
//		default:
//			time.Sleep(1 * time.Second)
//		}
//	}
//}
