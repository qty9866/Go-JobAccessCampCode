package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch1 = make(chan int)

func A() {
	fmt.Println("我是等待函数，我现在开始等")
	defer close(ch1)
	<-ch1
	fmt.Println("收到消息了，现在结束")
	wg.Done()
}
func B() {
	fmt.Println("我是发令员 5s后发送起跑信号")
	time.Sleep(5 * time.Second)
	ch1 <- 1
	wg.Done()
}

func main() {
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}
