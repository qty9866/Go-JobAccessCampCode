package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	println("Contex...")
	ctx, cf := context.WithCancel(context.Background())
	go testCancel(ctx)
	<-time.After(time.Second * 5)
	cf()
	time.Sleep(time.Millisecond)
	fmt.Println("main结束")
}

func testCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //在外界接收到cancel的时候，done会读到数据
			fmt.Println("主协程取消了，我是子协程，也同时结束了")
			return
		default:
			fmt.Println("我是子协程，我正在运行中...")
			time.Sleep(time.Second)
		}
	}
}
