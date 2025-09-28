package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Aditya7880900936/grpc-go/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println(message)
		time.Sleep(time.Second * 2)
	}
	log.Printf("Streaming finished")
}
