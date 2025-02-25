package main

import (
	"log"
	"time"

	pb "github.com/PoudelAmrit123/proto"
)

func (s *helloServer) SayHelloServerStreamning(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamningServer) error {

	log.Printf("got request with name : %v ", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil

}
