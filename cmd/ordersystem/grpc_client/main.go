package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cgalimberti/gocleanarc/20-CleanArch/internal/infra/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderServiceClient(conn)

	r, err := c.ListOrders(context.Background(), &pb.Blank{})
	//
	if err != nil {
		log.Fatalf("could not list orders: %v", err)
	}

	fmt.Printf("Orders: %v\n", r.Orders)
}
