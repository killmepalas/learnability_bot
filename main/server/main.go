package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	tgpb "learnability_bot/pb"
	"log"
	"net"
)

type server struct{}

func (s *server) Update(ctx context.Context, in *tgpb.UpdateRequest) (*tgpb.UpdateResponse, error) {

	if err := notify(in.CourseId, in.Course, in.Element, in.ElementType, in.ActionType); err != nil {
		log.Fatalf("Ошибка отправки %v", err)
		return &tgpb.UpdateResponse{
			Text: "don't send",
		}, nil
	} else {
		log.Printf("пять минут полёт нормальный")
		return &tgpb.UpdateResponse{
			Text: "succeed",
		}, nil
	}

}

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8090"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	tgpb.RegisterUpdateServiceServer(grpcServer, &server{})
	log.Println("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		return
	}

	//listener, err := net.Listen("tcp", "localhost:8090")
	//if err != nil {
	//	panic(err)
	//}
	//
	//s := grpc.NewServer()
	//tgpb.RegisterUpdateServiceServer(s, &server{})
	//if err := s.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
