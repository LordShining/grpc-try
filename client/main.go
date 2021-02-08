package main

import (
	"context"
	"fmt"

	//"os"
	"sync"
	"time"

	pb "github.com/LordShining/grpc-try/pb"
	go_uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

const (
	address = "localhost:36600"
)

func working(done func(), comments ...string) {
	defer done()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewBalancerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	r, err := c.Working(ctx, &pb.WorkRequest{Id: go_uuid.NewV4().String(), Comments: comments})
	if err != nil {
		fmt.Printf("could not doing: %v\n", err)
	}
	if r.Result {
		fmt.Printf("work id: %s finished\n", r.GetId())
	} else {
		fmt.Printf("work id: %s failed\n", r.GetId())
	}
}

func main() {
	fmt.Println("start client")

	newName := []string{"Bob", "Alice", "Green", "Charles", "Martin", "Joan", "Lily", "William"}
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		fmt.Printf("start task id: %d\n", i)
		wg.Add(1)
		go working(wg.Done, newName[i%len(newName)], "'s task")
		time.Sleep(time.Duration(2) * time.Second)
	}
	wg.Wait()
}
