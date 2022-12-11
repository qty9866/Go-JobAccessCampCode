package main

import "sync"

//  一般为了保证锁不出错，都内嵌在结构体中

type SafeCount struct {
	totalNum         int
	totalLetterCount int
	totalWordCount   int

	// ...
	sync.Mutex
}

var counter *SafeCount = &SafeCount{}

func (s *SafeCount) countNumber(totalNum, totalLetterCount, totalWordCount int) {
	s.Lock()
	defer s.Unlock()
	s.totalNum += s.totalNum
	//.....
}

func countNumber() {
	counter.countNumber(100, 5, 500)
}