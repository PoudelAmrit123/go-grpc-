package main

import (
	// "context"
	"io"
	"log"

	pb "github.com/PoudelAmrit123/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil

		}
		if err != nil {
			return nil
		}

		log.Printf("Got the message from the client : %v ", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

	}
}
