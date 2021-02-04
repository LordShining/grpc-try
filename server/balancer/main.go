package main

import (
	//"fmt"
	//"contex"
	//"sync"
	"context"
	"log"
	"net"

	pb "github.com/LordShining/grpc-try/pb"

	"google.golang.org/grpc"
)

const (
	//PORT ...
	PORT = ":66600"
)

//Server ...
type Server struct {
	pb.UnimplementedBalancerServer
}

//Request ...
type Request struct {
	id       int32
	comments []string
}

//注册服务
func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBalancerServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//GetWorks 任务接收与分配
func (s *Server) GetWorks(context.Context, *pb.WorkRequest) (*pb.Reply, error) {
	//to do
	return &pb.Reply{Id: 0, Result: true}, nil
}

//WorkerRegister 处理节点注册
func (s *Server) WorkerRegister(context.Context, *pb.WorkerRegisterRequest) (*pb.Reply, error) {
	return &pb.Reply{Id: 0, Result: true}, nil
}

//WorkerAlive 节点在线检测
func (s *Server) WorkerAlive(context.Context, *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Id: 0, Result: true}, nil
}
