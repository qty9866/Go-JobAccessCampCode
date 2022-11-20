package main

import (
	"fmt"
	"reflect"
)

// 接口里面只能有方法，不能放字段名之类的

func main() {
	var refrigerator Refrigerator
	fmt.Println(refrigerator.Size)
	var elephant Elephant
	fmt.Println(elephant.Name)

	/*var ip *int
	fmt.Println(ip)

	var putER PutElephantIntoRefrigerator // 默认是nil

	err := putER.OpenTheDoorOfRefrigerator(refrigerator)
	if err != nil {
		fmt.Printf("err:%v", err) // panic空指针
	}

	var players map[string]int 	// map可以直接初始化实体 但是接口不行
	players["age"] = 38*/

	// 注意：如果某个对象的成员方法有的在对象上有的在对象指针上，那么在给接口赋值时候，必须使用指针
	// 不然会报错
	var legend PutElephantIntoRefrigerator = &PutElephantIntoRefrigeratorImpl{}
	legend.OpenTheDoorOfRefrigerator(refrigerator)
	legend.PutElephantIntoRefrigerator(elephant, refrigerator)
	legend.CloseTheDoorOfRefrigerator(refrigerator)

	fmt.Println(reflect.TypeOf(funcTypeRealizeInterface)) //main.TestTypeImplInterface

	var i interface{}
	i = 35
	fmt.Println(reflect.TypeOf(i)) // int
	i = 3.14159265
	fmt.Println(reflect.TypeOf(i)) // float64
	//	 空接口在实际应用中非常有用例如fmt.Println()方法
	/*	func Println(a ...any) (n int, err error) {  a 是空接口类型
		return Fprintln(os.Stdout, a...)
	}*/

}

type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant, Refrigerator) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}

// TestTypeImplInterface 不管是什么类型 只要实现了接口的所有方法也就是实现了这个接口
type TestTypeImplInterface func()

// 实例化一个TestTypeImplInterface(func()类型)
var tt TestTypeImplInterface
var funcTypeRealizeInterface PutElephantIntoRefrigerator = tt

func (t TestTypeImplInterface) OpenTheDoorOfRefrigerator(_ Refrigerator) error {
	return nil
}
func (t TestTypeImplInterface) PutElephantIntoRefrigerator(_ Elephant, _ Refrigerator) error {
	return nil
}
func (t TestTypeImplInterface) CloseTheDoorOfRefrigerator(_ Refrigerator) error {
	return nil
}

type PutElephantIntoRefrigeratorImpl struct {
}

func (receiver *PutElephantIntoRefrigeratorImpl) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo something
	fmt.Println("打开冰箱门")
	return nil
}

func (receiver *PutElephantIntoRefrigeratorImpl) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	// todo
	fmt.Println("装进去")
	return nil
}

func (receiver *PutElephantIntoRefrigeratorImpl) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("关门")
	return nil
}

type Refrigerator struct {
	Size string
}

type Elephant struct {
	Name string
}

type Test1 interface {
	ABC() error
}

type Test2 interface {
	DEF() error
}

type RealizeTwo struct{}

func (RealizeTwo) ABC() error {
	fmt.Println("abc")
	return nil
}
func (RealizeTwo) DEF() error {
	fmt.Println("def")
	return nil
}

// 这个结构体同时可以实现两个接口
var i1 Test1 = RealizeTwo{}
var i2 Test2 = RealizeTwo{}

// Box1 同样接口也可以实现嵌套
type Box1 interface {
	Test1
	Test2
}
type Open interface {
	Open() error
}

type Close interface {
	Close() error
}

func (Refrigerator) Open() error {
	return nil
}

func (Refrigerator) Close() error {
	return nil
}

type Box interface {
	Open
	Close
}

var box Box1 = RealizeTwo{}
