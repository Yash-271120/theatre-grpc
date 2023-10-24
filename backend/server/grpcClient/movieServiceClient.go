package grpcclient

import (
	"log"
	"os"

	pb "github.com/Yash-271120/theatre-grpc/backend/server/proto"
)

type MovieRpcService struct {
	pb.MovieServiceClient
}

var MovieRpcServiceObj MovieRpcService

func InitConnectionToMovieService() {
	host, hostexists := os.LookupEnv("MOVIE_SERVICE_HOST")
	port, portexists := os.LookupEnv("MOVIE_SERVICE_PORT")

	if !hostexists {
		log.Println("host not fount using default - 127.0.0.1")
		host = "127.0.0.1"
	}

	if !portexists {
		log.Println("port not fount using default - 6666")
		port = "6666"

	}

	conn, err := NewGrpcClient(host, port)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	MovieRpcServiceObj.MovieServiceClient = pb.NewMovieServiceClient(conn)
}
