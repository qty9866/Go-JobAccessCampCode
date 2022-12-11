package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

/*
	文件目錄相關操作
*/
// 创建文件
func TestCreateFile(t *testing.T) {
	file, err := os.Create("test1.txt")
	if err != nil {
		fmt.Println("Create File err", err)
	} else {
		fmt.Println(file)
	}
}

// 创建目录
func TestCreateDir(t *testing.T) {
	err := os.Mkdir("newDir", os.ModePerm)
	if err != nil {
		println(err)
	}
}

// 删除文件
func TestRemoveFile(t *testing.T) {
	err := os.Remove("./test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 删除目录
func TestRemoveDir(t *testing.T) {
	err := os.RemoveAll("./newDir")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

// 获得工作目录
func TestGetWd(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wd)
}

// 更改目录
func TestChDir(t *testing.T) {
	err := os.Chdir("f:/GoProject")
	if err != nil {
		os.Exit(2)
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wd)
}

// 获得临时目录
func TestGetTempDir(t *testing.T) {
	dir := os.TempDir()
	fmt.Println(dir) //C:\Users\Hud\AppData\Local\Temp
}

// 文件重命名
func TestRename(t *testing.T) {
	err := os.Rename("./test1.txt", "./empty.txt")
	fmt.Println(err)
}

/*
文件的读写
*/
// 写文件
func TestFileWrite(t *testing.T) {
	err := os.WriteFile("./empty.txt", []byte(`Hello and my name is Hud!
And Nice to meet you!`), os.ModePerm)
	if err != nil {
		os.Exit(2)
	}
}

// 读文件
func TestReadFile(t *testing.T) {
	file, err := os.ReadFile("./empty.txt")
	if err != nil {
		os.Exit(2)
	}
	fmt.Println(string(file))
}

// 打开文件 设置权限
func TestOpenFile(t *testing.T) {
	// 文件只能读
	f, err := os.Open("empty.txt")
	if err != nil {
		fmt.Println("Open file error:", err)
		os.Exit(3)
	}
	fmt.Printf("%v\n", f.Name())
	// 根据第二个参数:设置读写或者没有该文件就进行创建、最高权限755
	f2, _ := os.OpenFile("./newDir/newCreated.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f2: %v\n", f2.Name())

	defer f.Close()
	defer f2.Close()
}

// 创建文件
func TestCreateFile1(t *testing.T) {
	//等价于 	os.OpenFile("name",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
	f, err := os.Create("test3.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("f.Name(): %v\n", f.Name())
	// 第一个参数 ，目录默认，Temp第二个参数 文件名前缀（看看下面创建出的临时文件）
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

// 读文件
func TestReadFile1(t *testing.T) {
	// 先打开文件
	f, err := os.Open("./newDir/newCreated.txt")
	if err != nil {
		fmt.Println("Open file error:", err)
		os.Exit(2)
	}
	// 创建缓冲区，读取文件内容
	buffer := make([]byte, 10)
	length, err := f.Read(buffer)
	if err != nil {
		fmt.Println("read file error", err)
	}
	fmt.Println("读到了的长度:", length)
	fmt.Println(string(buffer[:length]))
}

// 循环读文件
func TestForReadFile(t *testing.T) {
	f, _ := os.Open("./newDir/newCreated.txt")
	for {
		// 创建缓冲区 将文件读进缓冲区
		buffer := make([]byte, 10)
		n, err := f.Read(buffer)
		print(string(buffer[:n]))
		if err == io.EOF {
			break
		}
	}
}

// 文件写操作
func TestWriteFile(t *testing.T) {
	f, err := os.OpenFile("./newDir/newCreated.txt", os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("open file error:", err)
		os.Exit(2)
	}
	f.Write([]byte("I have finished my undergraduate education in Jiangsu Normal University."))
	f.Close()
}

// 写字符串
func TestWriteString(t *testing.T) {
	f, _ := os.OpenFile("./empty.txt", os.O_RDWR|os.O_TRUNC, 0777)
	n, err := f.WriteString("My name is Qty,hello~")
	fmt.Println("写入长度：", n)
	if err != nil {
		fmt.Println("write file error:")
	}
	f.Close()
}

// 指定位置开始写
func TestWriteInAssignedPos(t *testing.T) {
	f, _ := os.OpenFile("./empty.txt", os.O_RDWR, 0777)
	f.WriteAt([]byte("this is from WriteAt"), 22)
	f.Close()
	s := os.Environ()
	fmt.Printf("s: %v\n", s)
}
