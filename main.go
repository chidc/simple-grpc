package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	demo "simple-grpc/proto"
)

type server struct {
	demo.UnimplementedRestaurantLikeServiceServer
}
func (s *server) GetRestaurantLikeStat(ctx context.Context, req *demo.RestaurantLikeStatRequest) (*demo.RestaurantLikeStatResponse, error){
	return  &demo.RestaurantLikeStatResponse{
		Result: map[int32]int32{1:1,2:4},
	},nil
}

func main(){
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
	log.Println("Serving gRPC on 0.0.0.0:50051")
	log.Fatal(s.Serve(lis))
}
