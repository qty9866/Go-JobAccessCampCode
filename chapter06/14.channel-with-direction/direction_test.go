package main

import (
	"fmt"
	"testing"
)

func TestChannelWithDirection(t *testing.T) {
	c := make(chan int, 100)
	InOnly(c)
	OutOnly(c)
}

func InOnly(c chan<- int) {
	c <- 1
	// <-c    当c是单项入channel的时候，不能再从中取数字了
}

func OutOnly(c <-chan int) {
	o := <-c
	// c <- 1 当c是单项出channel的时候，不能再向channel中写数据了
	fmt.Println(o)
}
