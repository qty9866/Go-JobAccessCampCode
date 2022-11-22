package main

import (
	"fmt"
	"testing"
)

type Change interface {
	ChangeName(newName string)
	ChangeAge(newAge int)
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) ChangeName(newName string) {
	s.Name = newName
}
func (s Student) ChangeAge(newAge int) {
	s.Age = newAge
}

func TestVal(t *testing.T) {
	var stdChg Change
	stdChg = &Student{"QTY", 24}
	stdChg.ChangeAge(26)
	stdChg.ChangeName("Hud")
	fmt.Println(stdChg)
}

// TIPS: 这里注意一点
/*
	如果我们在定义方法的时候，只要有一个方法是定义在结构体指针上面的(也就是func (para *structName))
	那么，在对实现所有方法的接口变量进行实例化时，必须使用指针类型*Student，Student会报错没有实现该接口中所有的方法
*/

/*func TestStd(t *testing.T) {
	s := &Student{"Hud",24}
	s.ChangeName("QTY")
	fmt.Println("Now the student's name is:", s.Name)
}*/
