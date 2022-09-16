package mq

import (
	"errors"
	"sync"
	"time"
)

type Broker interface {
	publish(topic string, msg interface{}) error
	subscribe(topic string) (<-chan interface{}, error)
	unsubscribe(topic string, sub <-chan interface{}) error
	close()
	broadcast(msg interface{}, subscribers []chan interface{})
	setConditions(capacity int)
}

type BrokerImpl struct {
	exit     chan bool
	capacity int
	//
	topics map[string][]chan interface{}
	sync.RWMutex
}

func NewBroker() *BrokerImpl {
	return &BrokerImpl{
		exit:     make(chan bool),
		capacity: 24,
		topics:   make(map[string][]chan interface{}),
	}
}

func (bro *BrokerImpl) publish(topic string, msg interface{}) error {
	select {
	case <-bro.exit:
		return errors.New("broker closed")
	default:
		//bro.close()
	}
	//
	bro.RLock()
	subscribers, ok := bro.topics[topic]
	bro.RUnlock()
	if !ok {
		return nil
	}
	//
	bro.broadcast(msg, subscribers)
	return nil
}

func (bro *BrokerImpl) subscribe(topic string) (<-chan interface{}, error) {
	select {
	case <-bro.exit:
		return nil, errors.New("broker closed")
	default:
		//bro.close()
	}
	//
	ch := make(chan interface{}, bro.capacity)
	bro.Lock()
	bro.topics[topic] = append(bro.topics[topic], ch)
	bro.Unlock()
	return ch, nil
}
func (bro *BrokerImpl) unsubscribe(topic string, sub <-chan interface{}) error {
	select {
	case <-bro.exit:
		return errors.New("broker closed")
	default:
		//bro.close()
	}
	//
	bro.RLock()
	subscribers, ok := bro.topics[topic]
	bro.RUnlock()
	//
	if !ok {
		return nil
	}
	//
	var newSubs []chan interface{}
	for _, subscriber := range subscribers {
		if subscriber == sub {
			continue
		}
		newSubs = append(newSubs, subscriber)
	}
	//
	bro.Lock()
	bro.topics[topic] = newSubs
	bro.Unlock()
	return nil
}
func (bro *BrokerImpl) close() {
	select {
	case <-bro.exit:
		return
	default:
		close(bro.exit)
		bro.Lock()
		bro.topics = make(map[string][]chan interface{})
		bro.Unlock()
	}
	return
}
func (bro *BrokerImpl) broadcast(msg interface{}, subscribers []chan interface{}) {
	count := len(subscribers)
	concurrency := 1
	//
	switch {
	case count > 10000:
		concurrency = 4
	case count > 1000:
		concurrency = 3
	case count > 100:
		concurrency = 2
	default:
		concurrency = 1
	}
	pub := func(start int) {
		for j := start; j < count; j += concurrency {
			select {
			case subscribers[j] <- msg:
			case <-time.After(time.Millisecond * 5):
			case <-bro.exit:
				return
			}
		}
	}
	for i := 0; i < concurrency; i++ {
		go pub(i)
	}
}
func (bro *BrokerImpl) setConditions(capacity int) {
	bro.capacity = capacity
}
