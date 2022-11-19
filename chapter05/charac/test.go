package main

import "fmt"

func main() {
	// 切片定义方式复习
	// MAKE
	player1 := make(map[string]string)
	player1["name"] = "Lebron James"
	player1["age"] = "38"
	player1["team"] = "Los Angles Lakers"
	fmt.Println(player1)

	player2 := map[string]string{
		"name": "Derrick Rose",
		"age":  "34",
		"team": "Chicago Bulls",
	}
	fmt.Println(player2)
}
