package main

import (
	"context"
	"fmt"
	clientV3 "go.etcd.io/etcd/client/v3"
	"log"
	"strconv"
	"time"
)

func main() {
	etcdGetDemo()
	etcdWatchAndReactDemo()
}

func etcdGetDemo() {
	client, err := clientV3.New(clientV3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Get(context.TODO(), "a")
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range resp.Kvs {
		log.Printf("key:%s,value:%s", kv.Key, kv.Value)
	}
}

func etcdWatchAndReactDemo() {
	client, err := clientV3.New(clientV3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}

	dataCh := make(chan int)
	go func() {
		watcher := client.Watch(context.TODO(), "a")
		for respData := range watcher {
			whetherBreak := false
			evs := respData.Events
			for _, ev := range evs {
				i, err := strconv.Atoi(string(ev.Kv.Value))
				if err != nil {
					fmt.Println("不是数字,結束")
					whetherBreak = true
					break
				}
				dataCh <- i
			}
			if whetherBreak {
				break
			}
		}
	}()

	go func() {
		for i := range dataCh {
			_, err := client.Put(context.TODO(), "a", fmt.Sprintf("%d", i))
			if err != nil {
				fmt.Println("WARNING: 更新失败")
			}
		}
	}()

	time.Sleep(2 * time.Second)
}
