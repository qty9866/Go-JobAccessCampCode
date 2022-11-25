package main

import (
	"fmt"
	"sync"
)

// 闭包函数如果和goroutine一起使用，闭包里的参数值会因为共享变量而出现问题
// 解决方案：显式传入参数
func main() {
	iarr := []int{1, 2, 3, 4, 5, 6}
	wg := sync.WaitGroup{}
	wg.Add(len(iarr))
	for _, item := range iarr {
		go func(newItem int) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(newItem)
			}
		}(item)
	}
	wg.Wait()
}
