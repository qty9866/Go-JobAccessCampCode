package main

import "fmt"

func main() {
	var data string

	{
		var equipment IOInterface = &Soft{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Paper{}
		data = equipment.Read()
		fmt.Println(data)
	}
}

// IOInterface 定义一个带有Read方法的接口:IOInterface
type IOInterface interface {
	Read() string
}

type Soft struct{}

func (Soft) Read() string {
	return "lalalalalala软盘数据"
}

type Mag struct{}

func (Mag) Read() string {
	return "zizizizizizi~~~~"
}

type Paper struct{}

func (Paper) Read() string {
	return "从纸带读数据xxxx0000"
}
