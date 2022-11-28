package main

import (
	"fmt"
	"testing"
	"time"
)

func TestFinishWithChannelSignal(t *testing.T) {
	fmt.Println("Test Start")
	ch := make(chan bool)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("主协程通知我需要取消了")
				return
			default:
				fmt.Println("我是一个goroutine，我正在工作")
			}
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(3 * time.Second)
	ch <- false
}
