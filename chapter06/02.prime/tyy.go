package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	print("this is from function main")
	go goRoutine()
	wg.Wait()
}

func goRoutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("this is from goroutine", i)
	}
	defer wg.Done()
}
