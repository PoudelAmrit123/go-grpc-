package main

import (
	"context"

	pb "github.com/PoudelAmrit123/proto"
)

//SayHello(context.Context, *NoParam) (*HelloResponse, error)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello How are you",
	}, nil

}
