package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var qa = map[string]string{
	"你好":       "你好",
	"你是谁？":     "我是傻逼",
	"你在干嘛？":    "不是在跟你聊天呢吗？操",
	"你是男是女？":   "纯爷们",
	"今天天气怎么样？": "这个不知道，你自己看看呢",
	"再见":       "再见",
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("监听端口错误", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("建立连接失败:", err)
			continue
		}
		fmt.Println(conn)

		go talk(conn)
		//talk(conn)
	}
}

func talk(conn net.Conn) {
	defer fmt.Println("结束链接：", conn)
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				time.Sleep(time.Second)
				continue
			}
			log.Println("读取数据出错", err)
			continue
		}
		content := buf[:read]
		// 判断一下是否拿到了值(用ok)
		resp, ok := qa[string(content)]
		if !ok {
			resp = "对不起，这个问题老子还在学习中"
			log.Println("对不起，这个问题老子还在学习中")
		}
		_, err = conn.Write([]byte(resp))
		if err != nil {
			// todo handle error
			return
		}
		if string(content) == "再见" {
			break
		}
	}
}
