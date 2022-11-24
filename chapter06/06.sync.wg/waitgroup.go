package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	Name string
}

func (r Runner) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Println(r.Name, "开始跑步")
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	finish := time.Now()
	fmt.Println(r.Name, "跑完了", finish.Sub(start))
}

func main() {
	runnersCount := 10
	var runners []Runner

	wg := sync.WaitGroup{}
	wg.Add(runnersCount)

	for i := 0; i < runnersCount; i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("%d号选手", i),
		})
	}
	for _, runnerItem := range runners {
		go runnerItem.Run(&wg)
	}
	wg.Wait()
	fmt.Println("赛跑结束")
}
