package main

import (
	//"fmt"
	//"sync"
	"context"
	"log"
	"net"
	"strings"
	"time"

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
	portList  []string
	taskCount []int
	workerNum int
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

//Working 任务接收与分配
func (s *Server) Working(ctx context.Context, req *pb.WorkRequest) (*pb.Reply, error) {
	tWN := s.workerNum
	tPL := make([]string, tWN)
	copy(tPL, s.portList)
	tTC := make([]int, tWN)
	copy(tTC, s.taskCount)
	for i, v := range tPL {
		if tTC[(i-1)%tWN] > tTC[i] {
			//give task to worker i
			address := strings.Join([]string{"localhost"}, v)
			// Set up a connection to the server.
			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewBalancerClient(conn)

			// Contact the server and print out its response.
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
			defer cancel()
			r, err := c.Working(ctx, req)
			if err != nil {
				log.Fatalf("could not doing: %v", err)
			}
			return r, nil
		}
	}

	return &pb.Reply{Id: "0", Result: false}, nil
}

//WorkerRegister 处理节点注册
func (s *Server) WorkerRegister(ctx context.Context, req *pb.WorkerRegisterRequest) (*pb.Reply, error) {
	return &pb.Reply{Id: "0", Result: true}, nil
}

//WorkerAlive 节点在线检测
func (s *Server) WorkerAlive(ctx context.Context, req *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Id: "0", Result: true}, nil
}
