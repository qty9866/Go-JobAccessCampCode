package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	var input Interface = &fakeInterface{
		name:       "Hud",
		sex:        "男",
		baseWeight: 81,
		baseTall:   1.81,
		baseAge:    24,
	}

	for {
		func() {
			conn, err := net.Dial("tcp", "localhost:8080")
			if err != nil {
				log.Fatal("客户端端口监失败:", err)
			}
			defer conn.Close()
			fmt.Println("连接成功，开始发送数据")
			person, err := input.ReadPersonalInformation()
			if err != nil {
				log.Println("读取数据出错，请重新输入", err)
			}
			data, err := json.Marshal(person)
			if err != nil {
				log.Println("数据序列化出错", err)
				return
			}
			log.Println("讀取到的数据：", string(data))
			talk(conn, string(data))
		}()
		time.Sleep(time.Second)
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
