package main

import (
	"fmt"
	"sync"
)

// Queue 定义队列结构体
type Queue struct {
	data []interface{}
	sync.Mutex
}

// Push 队列添加数据,线程安全
func (q *Queue) Push(data interface{}) {
	q.Lock()
	q.data = append(q.data, data)
	q.Unlock()
}

// Pop 从队列中取数据
func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.data) > 0 {
		out := q.data[0]
		q.data = q.data[1:]
		return out, true
	}
	return nil, false
}

func main() {
	q := &Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
