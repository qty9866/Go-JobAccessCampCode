package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	maxNum := 200000
	result := make(chan int, maxNum/4)
	wg := sync.WaitGroup{}

	workerNum := 16
	wg.Add(workerNum)
	baseNumch := make(chan int, 10)
	for i := 0; i < workerNum; i++ {
		go func() {
			defer wg.Done()
			for onum := range baseNumch {
				if isPrime(onum) {
					result <- onum
				}
			}
		}()
	}

	for num := 0; num < maxNum; num++ {
		baseNumch <- num
	}
	close(baseNumch)
	wg.Wait()

	fmt.Println(len(result))
	finish := time.Now()
	fmt.Println(finish.Sub(startTime))
}

func isPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			return
		}
	}
	isPrime = true
	return
}
