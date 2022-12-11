package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "./empty.txt"
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File Open Error:", err)
		//return
		os.Exit(1) // 正常退出：0
	}
	var n int
	b := make([]byte, 20)
	for i := 0; i < 2; i++ {
		n, err = f.Read(b)
		if err != nil {
			fmt.Println("File Read Error:", err)
			os.Exit(2)
		}
	}
	defer f.Close()
	fmt.Println(string(b))
	fmt.Println("讀出的大小:", n)
}
