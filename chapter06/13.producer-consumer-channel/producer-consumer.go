package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Store struct {
	store chan int
	init  sync.Once
	Max   int
}

// 线程安全的初始化一个channel
func (s *Store) instrument() {
	s.init.Do(func() {
		s.store = make(chan int, s.Max)
	})
}

type Producer struct{}

func (Producer) Produce(s *Store) {
	fmt.Println("开始生产+1")
	s.store <- rand.Int()
}

type Consumer struct{}

func (c Consumer) Consume(s *Store) {
	fmt.Println("消费者消费-1", <-s.store)
}

func main() {
	s := &Store{
		Max: 10,
	}
	s.instrument()
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
