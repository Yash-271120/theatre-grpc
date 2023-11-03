package grpcclient

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/Yash-271120/theatre-grpc/backend/services/ticket/proto"
)

type HallRpcService struct {
	pb.HallServiceClient
}

var HallRpcServiceObj HallRpcService

func InitConnectionToHallService() {
	host, hostexists := os.LookupEnv("HALL_SERVICE_HOST")
	port, portexists := os.LookupEnv("HALL_SERVICE_PORT")

	if !hostexists {
		log.Println("host not fount using default - 127.0.0.1")
		host = "127.0.0.1"
	}

	if !portexists {
		log.Println("port not fount using default - 7777")
		port = "7777"

	}

	conn, err := NewGrpcClient(host, port)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("Connected to Hall Service on port %v:%v", host, port)
	HallRpcServiceObj.HallServiceClient = pb.NewHallServiceClient(conn)
}

func CallUpdateHall(req *pb.UpdateHallAvailableRequest, client *HallRpcService) (*pb.UpdateHallAvailableResponse, error) {
	ctx, cancel := context.WithTimeout((context.Background()), time.Second)
	defer cancel()
	response, err := client.UpdateHallAvailable(ctx, req)
	if err != nil {
		log.Fatalf("Error when calling UpdateHall: %s", err)
		return nil, err
	}
	return response, nil
}
