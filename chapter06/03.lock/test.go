package main

import "fmt"

func main() {
	test1()

}
func test1() {
	var players map[string]string
	players["James"] = "Lakers"
	fmt.Println(players) //panic: assignment to entry in nil map
	// 因为没有初始化，map不像array和基础类型在你定义就会给你初始化一个默认值
}
func test2() {
	var players map[string]string
	// 这里先进行实例化
	players = make(map[string]string)
	players["James"] = "Lakers"
	fmt.Println(players)
}
