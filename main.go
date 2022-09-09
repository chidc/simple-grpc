package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	demo "simple-grpc/proto/demo1"
)

type server struct {
	demo.UnimplementedRestaurantLikeServiceServer
}

func (s *server) GetRestaurantLikeStat(ctx context.Context, req *demo.RestaurantLikeStatRequest) (*demo.RestaurantLikeStatResponse, error) {
	return &demo.RestaurantLikeStatResponse{
		Result: map[int32]int32{1: 1, 2: 4},
	}, nil
}

func main() {
	// Create a listener on TCP port
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	demo.RegisterRestaurantLikeServiceServer(s, &server{})
	// Serve gRPC Server
	go func() {
		log.Println("Serving gRPC on 0.0.0.0:50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = demo.RegisterRestaurantLikeServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
}
