package main

import (
	"fmt"
	//"sync"
	"context"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/LordShining/grpc-try/pb"

	"google.golang.org/grpc"
)

//Server ...
type Server struct {
	pb.UnimplementedWorkerServer
	taskCount int
	port      string
	ch        chan int
}

const (
	address = "localhost:36600"
)

//注册服务
func main() {
	fmt.Println("start worker")
	if len(os.Args) != 2 {
		fmt.Println("need a port(like 34567)")
		return
	}
	port := strings.Join([]string{":", os.Args[1]}, "")
	fmt.Printf("try to register on port %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}
	//向balancer注册
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewBalancerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	r, err := c.WorkerRegister(ctx, &pb.WorkerRequest{Port: port})
	if err != nil {
		fmt.Printf("could not register: %v\n", err)
	}
	fmt.Println(r)

	s := grpc.NewServer()
	server := &Server{port: port, ch: make(chan int)}
	go keepAlive(server)
	pb.RegisterWorkerServer(s, server)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
	server.ch <- 0
}

//Working 处理任务
func (s *Server) Working(ctx context.Context, req *pb.WorkRequest) (*pb.Reply, error) {
	fmt.Printf("working on task: %s\n", req.GetId())
	time.Sleep(time.Duration(25) * time.Second)
	fmt.Println("Finished")
	return &pb.Reply{Id: req.GetId(), Result: true}, nil
}

//心跳请求
func keepAlive(s *Server) {
	exit := false
	go quitKeep(&exit, s.ch)
	for !exit {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			fmt.Printf("did not connect: %v\n", err)
		}
		defer conn.Close()
		c := pb.NewBalancerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		r, err := c.WorkerAlive(ctx, &pb.WorkerRequest{Port: s.port})
		if err != nil {
			fmt.Printf("could not register: %v\n", err)
		}
		if r.Result {
			//fmt.Println("alive")
			time.Sleep(time.Duration(9) * time.Second)
		}
	}
}

func quitKeep(exit *bool, ch chan int) {
	<-ch
	*exit = true
	fmt.Println("quit")
}
