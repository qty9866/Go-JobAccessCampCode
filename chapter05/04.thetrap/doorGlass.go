package main

import "fmt"

var _ Door = &GlassDoor{}

// 把一个GlassDoor对象的指针赋值给一个Door类型的变量，赋值到_(也就是nowhere)，强制去实现这个接口里面的所有方法

type GlassDoor struct{}

func (d *GlassDoor) Unlock() {
	fmt.Println("GlassDoor Unlock")
}

func (d *GlassDoor) Lock() {
	fmt.Println("GlassDoor Lock")
}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}
func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Open")
}
