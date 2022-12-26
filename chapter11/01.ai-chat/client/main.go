package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("客户端端口监失败:", err)
	}
	//talk(conn, "你好")
	//talk(conn, "你是谁？")
	//talk(conn, "今天天气怎么样？")
	//talk(conn, "再见")
	defer conn.Close()
	fmt.Println("连接成功，开始聊天")
	for {
		r := bufio.NewReader(os.Stdin)
		input, _, _ := r.ReadLine()
		if len(input) != 0 {
			talk(conn, string(input))
		}
	}
}

func talk(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Println("发送消息失败:", err)
	} else {
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if err != nil {
			log.Println("读取消息错误", err)
		} else {
			data := buf[:length]
			log.Println("去:", message, "回:", string(data))
		}
	}

}
