package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Aditya7880900936/grpc-go/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf(" Client Streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}

		log.Printf("Sent request with name: %v", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}
	log.Printf("LongGreet Response: %v", res)
}
