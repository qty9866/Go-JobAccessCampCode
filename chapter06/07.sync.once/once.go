package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Rank struct {
	standard []string
}

var globalRank = &Rank{}
var once sync.Once
var wg sync.WaitGroup

/*func init() {
	globalRank.standard = []string{"Aisa"}
}*/

func initGlobalRankStandard(standard []string) {
	// 即使每个协程都会调用，但是只会执行一次
	// 多线程中共享内容的初始化，且只初始化一次
	once.Do(func() {
		globalRank.standard = standard
	})
}

func main() {
	wg.Add(10)
	standard := []string{"Asia"}
	for i := 0; i < 10; i++ {
		go func() {
			initGlobalRankStandard(standard)
			fmt.Println(rand.Int())
			wg.Done()
		}()
	}
	fmt.Println("Work is Done,Rank:", globalRank.standard)
}
