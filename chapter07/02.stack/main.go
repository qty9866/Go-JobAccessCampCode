package main

import "fmt"

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.data = append([]interface{}{data}, s.data...)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) > 0 {
		out := s.data[0]
		s.data = s.data[1:]
		return out, true
	} else {
		return nil, false
	}
}

func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Pop()) // 4 true
	fmt.Println(s.Pop()) // 3 true
	fmt.Println(s.Pop()) // 2 true
	fmt.Println(s.Pop()) // 1 true
	fmt.Println(s.Pop()) // nil false
}
