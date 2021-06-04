package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/paedayz/go-grpc-course/greet/greetpb"
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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("GreetManyTimes function was invoked with a streaming request\n")
	result := ""
	for {
		req, err := stream.Recv()
		fmt.Print("Server Recieve Data\n")
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
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
