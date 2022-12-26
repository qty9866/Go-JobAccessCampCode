package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	r := strings.NewReader("hello world from hud!")
	r2 := bufio.NewReader(r)    //创建一个新的reader，将r进行封装
	s, _ := r2.ReadString('\n') //以换行作为结尾
	fmt.Printf("s: %v\n", s)
}

// Reset：丢弃缓冲中的数据，清除任何错误，将b重设为其下层从r读取数据
func TestReset(t *testing.T) {
	s1 := strings.NewReader("ABCDEF")
	s2 := strings.NewReader("123456")
	reader := bufio.NewReader(s1) //s1读入缓冲区
	readString, _ := reader.ReadString('\n')
	fmt.Println(readString)
	reader.Reset(s2)
	s, _ := reader.ReadString('5')
	fmt.Println(s)
}
