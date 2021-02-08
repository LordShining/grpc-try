package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	pb "github.com/LordShining/grpc-try/pb"

	"google.golang.org/grpc"
)

const (
	port = ":36600"
)

//Server ...
type Server struct {
	pb.UnimplementedBalancerServer
	//nodeList  []*WorkerNode
	workerMap  map[string]*WorkerNode
	workerNum  int
	minTaskNum int
	mu         sync.Mutex
	//ch         chan int
}

//WorkerNode ...
type WorkerNode struct {
	//port      string
	taskCount int
	alive     bool
	timer     *time.Timer
}

//注册服务
func main() {
	fmt.Println("start balancer")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterBalancerServer(s, &Server{ /*ch: make(chan int),*/ workerMap: make(map[string]*WorkerNode)})

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}

//Working 任务接收与分配
func (s *Server) Working(ctx context.Context, req *pb.WorkRequest) (*pb.Reply, error) {
	//var wg sync.WaitGroup
	//wg.Add(1)
	s.mu.Lock()
	tm := make(map[string]*WorkerNode)
	for k, v := range s.workerMap {
		tm[k] = v
	}

	s.mu.Unlock()
	id := req.GetId()
	fmt.Printf("get task %s\n", id)
	for k, v := range tm {
		fmt.Println(k, v)
		//if tWN == 1 || tTC[(i-1+tWN)%tWN] >= tTC[i] {
		if v.alive && v.taskCount <= s.minTaskNum {
			v.taskCount++
			s.mu.Lock()
			s.minTaskNum = v.taskCount
			s.mu.Unlock()
			//defer wg.Done()
			//give task to worker i
			fmt.Printf("give task %s to worker %s\n", id, k)
			address := strings.Join([]string{"localhost", k}, "")
			// Set up a connection to the server.

			conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				fmt.Printf("did not connect: %v\n", err)
			}
			defer conn.Close()
			c := pb.NewWorkerClient(conn)

			// Contact the server and print out its response.
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
			defer cancel()
			r, err := c.Working(ctx, req)
			if err != nil {
				fmt.Printf("could not doing: %v\n", err)
			}
			return r, nil
		}
	}

	//wg.Wait()
	return &pb.Reply{Id: "0", Result: false}, nil
}

//WorkerRegister 处理节点注册
func (s *Server) WorkerRegister(ctx context.Context, req *pb.WorkerRequest) (*pb.Reply, error) {
	port := req.GetPort()
	if v, ok := s.workerMap[port]; ok { //已经存在，则刷新生存时间
		v.timer.Reset(time.Duration(10) * time.Second)
		if !v.alive {
			v.alive = true
			go workerDied(port, v, s)
			s.mu.Lock()
			s.workerNum++
			tempNum := s.workerNum
			tempMap := make(map[string]*WorkerNode)
			for k, v := range s.workerMap {
				tempMap[k] = v
			}
			s.mu.Unlock()
			fmt.Printf("worker %s back to line\n%d worker online\n", port, tempNum)
			for k, v := range tempMap {
				fmt.Printf("worker %s online:%v\n", k, v.alive)
			}
		}
		return &pb.Reply{Id: "0", Result: true}, nil
	}
	temp := WorkerNode{
		alive: true,
		timer: time.NewTimer(time.Duration(10) * time.Second),
	}
	go workerDied(port, &temp, s)
	s.mu.Lock()
	s.workerMap[port] = &temp
	s.workerNum++
	tn := s.workerNum
	tm := make(map[string]*WorkerNode)
	for k, v := range s.workerMap {
		tm[k] = v
	}
	s.mu.Unlock()

	fmt.Printf("new worker register success\non port %s\n", port)
	fmt.Printf("%d worker online\n", tn)
	for k, v := range tm {
		fmt.Printf("worker %s online:%v\n", k, v.alive)
	}
	return &pb.Reply{Id: "0", Result: true}, nil
}

//WorkerAlive 节点在线检测
func (s *Server) WorkerAlive(ctx context.Context, req *pb.WorkerRequest) (*pb.Reply, error) {
	port := req.GetPort()
	temp := s.workerMap[port]
	temp.timer.Reset(time.Duration(10) * time.Second)
	if !temp.alive {
		temp.alive = true
		go workerDied(port, temp, s)
		s.mu.Lock()
		s.workerNum++
		tempNum := s.workerNum
		tempMap := make(map[string]*WorkerNode)
		for k, v := range s.workerMap {
			tempMap[k] = v
		}
		s.mu.Unlock()
		fmt.Printf("worker %s back to line\n%d worker online\n", port, tempNum)
		for k, v := range tempMap {
			fmt.Printf("worker %s online:%v\n", k, v.alive)
		}
	}
	return &pb.Reply{Id: "0", Result: true}, nil
}

//超时触发
func workerDied(port string, wn *WorkerNode, s *Server) {
	<-wn.timer.C
	wn.alive = false
	s.mu.Lock()
	s.workerNum--
	temp := s.workerNum
	s.mu.Unlock()
	fmt.Printf("worker %s offline\n%d worker online\n", port, temp)
}
