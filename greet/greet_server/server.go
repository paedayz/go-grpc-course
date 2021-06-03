package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

// All server type will have Greet() interface function
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()

	// create GreetService
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fial to serve: %v", err)
	}
}
