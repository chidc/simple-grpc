package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	demo "simple-grpc/proto/demo1"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()
	client := demo.NewRestaurantLikeServiceClient(cc)
	request := &demo.RestaurantLikeStatRequest{ResIds: []int32{1, 2, 3}}

	resp, _ := client.GetRestaurantLikeStat(context.Background(), request)
	fmt.Println("receive response => [%v]\n", resp.Result)
}
