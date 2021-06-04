package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/paedayz/go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	// connect with service
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	// close connection when all of main code is run success
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// doUnary(c)

	// doServerStreaming(c)

	doClientStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "pae",
			LastName:  "phasit",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "pae",
			LastName:  "pad",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet many time RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pae",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pad",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "P'Por",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "P'Plam",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mon",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	// we iterate over our slice and send each message indivdually
	for _, req := range requests {
		fmt.Printf("Sending req : %v\n", req)
		stream.Send(req)
		time.Sleep(300 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}

	fmt.Printf("LongGreet Response: %v", res)
}
