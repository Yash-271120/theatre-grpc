package main

import (
	"log"
	"net"
	"os"

	grpcclient "github.com/Yash-271120/theatre-grpc/backend/services/ticket/grpcClient"
	grpcserver "github.com/Yash-271120/theatre-grpc/backend/services/ticket/grpcServer"
	"github.com/Yash-271120/theatre-grpc/backend/services/ticket/models"
	pb "github.com/Yash-271120/theatre-grpc/backend/services/ticket/proto"
	env "github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func initEnv() {
	if err := env.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println("Environment variables loaded successfully")
}

func main() {
	initEnv()
	host, hostexists := os.LookupEnv("HOST")
	port, portexists := os.LookupEnv("PORT")

	if !hostexists {
		log.Println("host not fount using default - 127.0.0.1")
		host = "127.0.0.1"
	}
	if !portexists {
		log.Println("port not fount using default - 8888")
		port = "8888"
	}
	// create a listener on TCP port 8888
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v:%v : %v", host, port, err)
	}

	// create a gRPC server object
	s := grpc.NewServer()

	// initialize the hall service client
	grpcclient.InitConnectionToHallService()

	// attach the Ping service to the server
	pb.RegisterTicketServiceServer(s, &grpcserver.TicketRPCServer{})

	// connect to database
	models.InitDB()

	log.Printf("Starting gRPC server on port %v:%v...", host, port)
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v:%v : %v", host, port, err)
	}
}
