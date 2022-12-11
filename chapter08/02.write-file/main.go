package main

import (
	"fmt"
	"os"
)

func main() {
	filepath := "F:/Learning-JobAccess-Camp/chapter08/02.write-file/new"
	f, err := os.OpenFile(filepath, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Open file error:", err)
		os.Exit(1)
	}

	defer f.Close()
	//buf := make([]byte,10)
	//var n int
	_, err = f.Write([]byte("this is line1\n"))
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte("this is line2\n"))
	f.WriteAt([]byte("CHANGED"), 0)

}
