package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 线程安全的map,但是跟一般的map是完全不一样的
	m := sync.Map{} // 虽然名字叫map实际上里面是个struct
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Store(i, 1)
			for {
				// 不能直接通过下标，要读(Load)
				v, ok := m.Load(i)
				if !ok {
					continue
				}
				// 不能直接写值，要写(Store)
				m.Store(i, v.(int)+1)
				fmt.Println("i=", v)
			}
		}(i)
	}
	time.Sleep(4 * time.Second)
}
