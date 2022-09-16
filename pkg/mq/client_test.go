package mq

import (
	"fmt"
	"sync"
	"testing"
)

const topic = "JULIA"

func TestClient(t *testing.T) {
	b := NewClient()
	b.SetConditions(2)
	var wg sync.WaitGroup
	//
	for i := 0; i < 2; i++ {
		topic := fmt.Sprintf("Golang梦工厂%d", i)
		payLoad := fmt.Sprintf("JavaScript%d", i)
		//
		ch, err := b.Subscribe(topic)
		if err != nil {
			t.Fatal(err)
		}
		//
		wg.Add(1)
		go func() {
			e := b.GetPayLoad(ch)
			fmt.Println(e)
			if e != payLoad {
				t.Fatalf("%s expected %s but get %s", topic, payLoad, e)
			}
			if err := b.UnSubscribe(topic, ch); err != nil {
				t.Fatal(err)
			}
			wg.Done()
		}()
		//
		if err := b.Publish(topic, payLoad); err != nil {
			t.Fatal(err)
		}
	}
}

//func TestOnceTopic(t *testing.T) {
//	m := NewClient()
//	m.SetConditions(2)
//	ch, err := m.Subscribe(topic)
//	if err != nil {
//		t.Fatal("Subscribe failed")
//		return
//	}
//	go OncePub(m)
//	OnceSub(ch, m)
//	defer m.Close()
//}
//
//func OncePub(c *Client) {
//	t := time.NewTicker(2 * time.Second)
//	defer t.Stop()
//	for {
//		select {
//		case <-t.C:
//			err := c.Publish(topic, "julia")
//			if err != nil {
//				fmt.Println("Publish failed")
//			}
//		default:
//
//		}
//	}
//}
//
//func OnceSub(m <-chan interface{}, c *Client) {
//	for {
//		val := c.GetPayLoad(m)
//		fmt.Printf("get message is %s\n", val)
//	}
//}
