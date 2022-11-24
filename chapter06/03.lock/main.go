package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//countDict()
	//countDictGoroutineSafe()
	//countDictForgetUnlock()
	//countDictLockPrice()
	countDictGoroutineSafeRW()
}

func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题
		}()
	}
	wg.Wait()
	fmt.Println("一共有", totalCount, "字")
}

func countDictGoroutineSafe() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			totalCountLock.Lock()
			totalCount += 100
			totalCountLock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("一共有", totalCount, "字")
}

func countDictLockPrice() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func(pageNum int) {
			defer wg.Done()
			totalCountLock.Lock()
			totalCount += 100
			if pageNum == 3 {
				time.Sleep(3 * time.Second)
			}
			totalCountLock.Unlock()
		}(p)
	}
	wg.Wait()
	fmt.Println("一共有", totalCount, "字")
}
func countDictForgetUnlock() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			// fmt.Print("正在统计第", p, "页，")
			totalCountLock.Lock()
			totalCount += 100 // totalCount = totalCount + 100 // 注意，这里有重复覆盖的问题
			// totalCountLock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictGoroutineSafeRW() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	totalCountLock := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(5)

	go func() {
		result := map[int64]struct{}{}
		for {
			totalCountLock.Lock()
			fmt.Println(totalCount)
			result[totalCount] = struct{}{}
			totalCountLock.Unlock()
		}
		fmt.Println("result", result)
	}()
	for p := 0; p < 5; p++ {
		go func() {
			defer wg.Done()
			fmt.Println("写锁开始时间：")
			totalCountLock.Lock()
			fmt.Println("写锁拿到锁时间")
			totalCount += 100
			totalCountLock.Unlock()

		}()
	}
	wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("一共有", totalCount, "字")
	fmt.Println()
}
