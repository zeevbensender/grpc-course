package main

import (
	"context"
	"log"

	pb "github.com/zeevbensender/grpc-clement-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Lupo",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", res.Result)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
