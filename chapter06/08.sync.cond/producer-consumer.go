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
	pCond     *sync.Cond
	cCond     *sync.Cond
}

type Producer struct{}

func (Producer) Produce(store *Store) {
	store.lock.Lock()
	defer store.lock.Unlock()
	if store.DataCount == store.Max {
		fmt.Println("生产者在等出货")
		store.pCond.Wait()
	}
	fmt.Println("生产+1")
	store.DataCount++
	store.cCond.Signal()
}

type Consumer struct{}

func (c Consumer) Consume(store *Store) {
	store.lock.Lock()
	defer store.lock.Unlock()
	if store.DataCount == 0 {
		fmt.Println("没货了，消费者在等")
		store.cCond.Wait()
	}
	fmt.Println("消费-1")
	store.DataCount--
	store.pCond.Signal()
}

func main() {
	s := &Store{
		Max: 10,
	}
	s.pCond = sync.NewCond(&s.lock)
	s.cCond = sync.NewCond(&s.lock)
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
