package main

import (
	"context"

	pb "github.com/Aditya7880900936/grpc-go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello I am Aditya and It is a Unary RPC call",
	}, nil
}
