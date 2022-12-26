package main

import (
	"Learning-JobAccess-Camp/pkg/apis"
	"encoding/json"
	"flag"
	"fmt"
	goBMI "github.com/armstrongli/go-bmi"
	"log"
	"net"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("监听端口错误", err)
	}

	rank := NewFatRateRank()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("建立连接失败:", err)
			continue
		}
		fmt.Println(conn)

		go talk(conn, rank)
	}
}

func talk(conn net.Conn, rank *FatRateRank) {
	defer fmt.Println("结束链接：", conn)
	defer conn.Close()
	for {
		finalReceivedMes := make([]byte, 0, 1024)
		EncounterError := false
		for {
			buf := make([]byte, 1024)
			read, err := conn.Read(buf)
			if err != nil {
				log.Println("读取数据出错", err)
				log.Println("重新读取", err)
				EncounterError = true
				time.Sleep(time.Second)
				break
			}
			if read == 0 {
				break
			}
			readContent := buf[:read]
			finalReceivedMes = append(finalReceivedMes, readContent...)
			if read < len(buf) {
				break
			}
		}
		if EncounterError {
			conn.Write([]byte("服务器获取数据失败，请重新输入"))
			continue
		}

		pi := &apis.PersonalInformation{}
		if err := json.Unmarshal(finalReceivedMes, pi); err != nil {
			conn.Write([]byte("输入的数据无法解析，请重新输入"))
			log.Println("输入的数据无法解析，请重新输入", err)
			continue
		}

		bmi, err := goBMI.BMI(float64(pi.Weight), float64(pi.Tall))
		if err != nil {
			conn.Write([]byte("无法计算BMI，请重新输入"))
			continue
		}
		rate := goBMI.CalcFatRate(bmi, int(pi.Age), pi.Sex)

		rank.inputRecord(pi.Name, rate)
		rankId, _ := rank.getRank(pi.Name)

		conn.Write([]byte(fmt.Sprintf("姓名：%s, BMI：%v，体脂率：%v，排名：%d", pi.Name, bmi, rate, rankId)))
		break
	}
}
