package main

import (
	"fmt"
	"sort"
)

type Button struct {
	Floor int
}

type Elevator struct {
	buttons  Buttons
	position int
}

type Buttons []*Button

// 实现buttons缺少的方法

func (b Buttons) Len() int {
	return len(b)
}

func (b Buttons) Less(i, j int) bool {
	return b[i].Floor < b[j].Floor
}

func (b Buttons) Swap(i, j int) {
	b[i].Floor, b[j].Floor = b[j].Floor, b[i].Floor
}

func main() {
	ev := Elevator{
		position: 2,
		buttons: []*Button{
			{Floor: 3},
			{Floor: 1},
			{Floor: 5},
			{Floor: 2},
			{Floor: 4},
		},
	}
	sort.Sort(ev.buttons)
	//fmt.Printf("%#v", ev.buttons)
	for _, item := range ev.buttons {
		fmt.Println(item.Floor)
	}
}
