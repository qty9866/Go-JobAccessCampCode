package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
}

type Producer struct{}

func (Producer) Produce(store *Store) {
	store.lock.Lock()
	defer store.lock.Unlock()
	if store.DataCount == store.Max {
		fmt.Println("Producer:'Store is full,stop Producing' ")
		return
	}
	fmt.Println("生产+1")
	store.DataCount++
}

type Consumer struct{}

func (c Consumer) Consume(store *Store) {
	store.lock.Lock()
	defer store.lock.Unlock()
	if store.DataCount == 0 {
		fmt.Println("Consumer:'None Data,stop consuming'")
		return
	}
	fmt.Println("消费者消费-1")
	store.DataCount--
}

func main() {
	s := &Store{
		Max: 10,
	}
	prodCount, consCount := 50, 50
	for i := 0; i < prodCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < consCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Consumer{}.Consume(s)
			}
		}()
	}
	time.Sleep(1 * time.Second)
}
