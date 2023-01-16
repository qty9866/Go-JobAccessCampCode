package main

import (
	"Learning-JobAccess-Camp/pkg/apis"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect:%v", err)
	}
	c := apis.NewRankServiceClient(conn)
	ret, err := c.Register(context.TODO(), &apis.PersonalInformation{
		Id:     1,
		Name:   "James",
		Sex:    "男",
		Tall:   2.03,
		Weight: 103,
		Age:    37,
	})
	if err != nil {
		log.Fatal("注册失败", err)
	}
	log.Println("注册成功", ret)

	log.Println("开始批量注册")
	regCli, err := c.RegisterPersons(context.TODO())
	if err != nil {
		log.Fatal("獲取批量注册客户端失败", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{
		Name:   "hud-3",
		Sex:    "男",
		Tall:   1.80,
		Weight: 81,
		Age:    24,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{
		Name:   "hud-1",
		Sex:    "男",
		Tall:   1.805,
		Weight: 79,
		Age:    24,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{
		Name:   "hud-2",
		Sex:    "男",
		Tall:   1.81,
		Weight: 81,
		Age:    24,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	resp, err := regCli.CloseAndRecv()
	if err != nil {
		log.Fatal("无法接收结果")
	}

	log.Println("批量注册结果", resp.String())
}
