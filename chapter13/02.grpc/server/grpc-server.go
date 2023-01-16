package main

import (
	"Learning-JobAccess-Camp/pkg/apis"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	startGRPCServer(ctx)
}

func startGRPCServer(ctx context.Context) {
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer([]grpc.ServerOption{}...)
	apis.RegisterRankServiceServer(server, &rankServer{
		persons:  map[string]*apis.PersonalInformation{},
		personCh: make(chan *apis.PersonalInformation),
	})
	go func() {
		select {
		case <-ctx.Done():
			server.Stop()
		}
	}()
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to server :%v", err)
	}

}

//func startGRPCServer(ctx context.Context) {
//	listen, err := net.Listen("tcp", "0.0.0.0:9090")
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer([]grpc.ServerOption{}...)
//	apis.RegisterRankServiceServer(s, &rankServer{
//		persons:  map[string]*apis.PersonalInformation{},
//		personCh: make(chan *apis.PersonalInformation, 1024),
//	})
//	go func() {
//		select {
//		case <-ctx.Done():
//			s.Stop()
//		}
//	}()
//	if err := s.Serve(listen); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
