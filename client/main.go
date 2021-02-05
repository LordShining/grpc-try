package main

import (
	"context"
	"log"

	//"os"
	"sync"
	"time"

	pb "github.com/LordShining/grpc-try/pb"
	go_uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

const (
	address = "localhost:66600"
)

func working(done func(), comments ...string) {
	defer done()

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
	r, err := c.Working(ctx, &pb.WorkRequest{Id: go_uuid.NewV4().String(), Comments: comments})
	if err != nil {
		log.Fatalf("could not doing: %v", err)
	}
	if r.Result {
		log.Printf("work id: %s finished", r.GetId())
	} else {
		log.Printf("work id: %s failed", r.GetId())
	}
}

func main() {

	newName := []string{"Bob", "Alice", "Green", "Charles", "Martin", "Joan", "Lily", "William"}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go working(wg.Done, newName[i%len(newName)], "'s task")
		time.Sleep(time.Duration(10) * time.Second)
	}
}
