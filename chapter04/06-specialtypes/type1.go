package main

import (
	"fmt"
	"math/rand"
)

type Math = int
type English = int
type Chinese = int

func main() {
	var mathScore1 = 148
	var mathScore2 Math = 150
	total := mathScore1 + mathScore2
	fmt.Println(total) // 说明是同一类型
	GetScoreOfStudent("")

	vg := &voteGame{
		students: []*Student{
			{name: fmt.Sprintf("%d", 1)},
			{name: fmt.Sprintf("%d", 2)},
			{name: fmt.Sprintf("%d", 3)},
			{name: fmt.Sprintf("%d", 4)},
			{name: fmt.Sprintf("%d", 5)},
			{name: fmt.Sprintf("%d", 6)},
			{name: fmt.Sprintf("%d", 8)},
			{name: fmt.Sprintf("%d", 9)},
			{name: fmt.Sprintf("%d", 10)},
			{name: fmt.Sprintf("%d", 11)},
			{name: fmt.Sprintf("%d", 12)},
			{name: fmt.Sprintf("%d", 13)},
		}}
	leader := vg.goRun()
	fmt.Println(leader)
	leader.Distribute()

	//	如果我现在实例化一个Student，虽然看上去他和Monitor一样，但是却无法直接进行赋值，因为类型不同，需要进行强转换
	var Hud = &Student{name: "QiTianyu"}
	var newMonitor Monitor = Monitor(*Hud)
	newMonitor.Distribute()

}

func GetScoreOfStudent(name string) (Math, English, Chinese) {
	//TODO
	return 0, 0, 0
}

type voteGame struct {
	students []*Student
}
type Leader = Student

func (v *voteGame) goRun() *Monitor {
	// TODO
	for _, item := range v.students {
		randInt := rand.Int()
		if randInt%2 == 0 {
			item.voteA(v.students[randInt%len(v.students)]) // TODO 用随机数来代替
		} else {
			item.voteD(v.students[randInt%len(v.students)])
		}

	}
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range v.students {
		if maxScore < item.agree {
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 {
		return (*Monitor)(v.students[maxScoreIndex]) // 直接将类型强制转换为Monitor
	}
	return nil
}

// Monitor 使用嵌套对象定义(继承)方式来定义班长
type Monitor Student

/*// Monitor 使用类型重定义
type Monitor Student*/

// Distribute 等于新定义了一个类型，而且拥有自己的方法
func (m *Monitor) Distribute() {
	fmt.Println("发作业了")
}

type FooooTestFuncRedefine func()

func (f *FooooTestFuncRedefine) test11() {

}

type Student struct {
	name     string
	agree    int
	disagree int
}

func (std *Student) voteA(target *Student) {
	std.agree++
}
func (std *Student) voteD(target *Student) {
	std.disagree++
}
