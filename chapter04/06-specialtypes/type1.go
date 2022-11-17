package main

import (
	"fmt"
	"math/rand"
)

type Math = int
type English = int
type Chinese = int

func main() {
	var mathScore1 int = 148
	var mathScore2 Math = 150
	total := mathScore1 + mathScore2
	fmt.Println(total) // 说明是同一类型
	GetScoreOfStudent("")

	vg := &voteGame{
		students: []*Student{
			&Student{name: fmt.Sprintf("%d", 1)},
			&Student{name: fmt.Sprintf("%d", 2)},
			&Student{name: fmt.Sprintf("%d", 3)},
			&Student{name: fmt.Sprintf("%d", 4)},
			&Student{name: fmt.Sprintf("%d", 5)},
			&Student{name: fmt.Sprintf("%d", 6)},
			&Student{name: fmt.Sprintf("%d", 8)},
			&Student{name: fmt.Sprintf("%d", 9)},
			&Student{name: fmt.Sprintf("%d", 10)},
			&Student{name: fmt.Sprintf("%d", 11)},
			&Student{name: fmt.Sprintf("%d", 12)},
			&Student{name: fmt.Sprintf("%d", 13)},
		}}
	leader := vg.goRun()
	fmt.Println(leader)
}

func GetScoreOfStudent(name string) (Math, English, Chinese) {
	//TODO
	return 0, 0, 0
}

type voteGame struct {
	students []*Student
}
type Leader = Student

func (v *voteGame) goRun() *Leader {
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
		return v.students[maxScoreIndex]
	}
	return nil
}

// Monitor 使用嵌套对象定义(继承)方式来定义班长
type Monitor struct {
	Student
}

/*// Monitor 使用类型重定义
type Monitor Student*/

func (m *Monitor) Distribute() {
	fmt.Println("发作业了")
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
