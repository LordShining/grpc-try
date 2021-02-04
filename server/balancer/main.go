package main

import (
	//"fmt"
	//"contex"
	//"sync"
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/LordShining/grpc-try/pb"
	"google.golang.org/grpc"
)

const (
	PORT = ":66600"
)

type Server struct {
	pb.UnimplementedBalancerServer
}
type Request struct {
	id       int32
	comments []string
}

//新请求监听
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//新服务器注册监听

func (s *Server) DoingWork(ctx context.Context, request *pb.WorkRequest) (*pb.WorkReply, error) {
	var r Request
	r.id = request.GetId()
	r.comments = request.GetComments()
	fmt.Print("Doing work: %d\n", r.id)
	for _, v := range r.comments {
		fmt.Println(v)
	}
	return &pb.WorkReply{Id: r.id, Result: true}, nil
}
