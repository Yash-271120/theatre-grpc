package grpcclient

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClient(address, port string) (*grpc.ClientConn, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	return conn, nil
}
