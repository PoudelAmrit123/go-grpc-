package main

import (
	// "context"
	// "io"
	// "log"

	"context"
	"log"
	"time"

	pb "github.com/PoudelAmrit123/proto"
)

func CallSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {

	log.Printf("Client Streamng started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names :%v", err)
	}

	for _, names := range names.Names {

		req := &pb.HelloRequest{
			Name: names,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with the name : %s", names)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Cient streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving : %v", err)
	}

	log.Printf("%v", res.Message)

}

// func (s *helloServer) SayHelloClientStreamning(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamningServer) error {

// 	log.Printf("received the req %v", req.Names)

// 	for _, names := range req.Names {
// 		res := &pb.HelloResponse{
// 			Message: "Hello" + names,
// 		}
// 		if err := stream.Send(res); err != nil {
// 			log.Printf("Error while streaming the file ", err)
// 		}
// 	}
// 	return nil

// }
