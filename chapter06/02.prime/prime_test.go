package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRunPrime(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	for num := 2; num < 200000; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("总共耗时", finishTime.Sub(startTime)) //总共耗时 3.8353748s
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func TestRunPrime2(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println("第一个协程开始工作")
		defer wg.Done()
		result = append(result, collectPrime(2, 100000)...) // 这里使用...是将[]int转成不定长的参数
		fmt.Println("第一个协程完成工作")
	}()
	go func() {
		fmt.Println("第二个协程开始工作")
		defer wg.Done()
		result = append(result, collectPrime(100000, 200000)...)
		fmt.Println("第二个协程完成工作")
	}()
	wg.Wait()
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("总共耗时", finishTime.Sub(startTime)) //总共耗时 3.8353748s
}

func collectPrime(start, end int) (result []int) {
	for num := start; num <= end; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	return
}
