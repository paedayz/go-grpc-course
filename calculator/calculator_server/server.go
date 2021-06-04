package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/paedayz/go-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber

	// created *SumResponse
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}

	return res, nil

}

func main() {
	fmt.Println("Caculator Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()

	// create GreetService
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fial to serve: %v", err)
	}
}
