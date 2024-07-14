package main

import (
	// "context"
	"io"
	"log"

	pb "github.com/PoudelAmrit123/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			return err
		}
		// log.Printf("Got the request with the name %v", req)
		messages = append(messages, req.Name)
		log.Println(messages)

	}

}

// func CallSayHelloClientStreaming(client pb.GreetServiceClient, name *pb.NameList) {

// 	log.Printf("Streaming started")

// 	stream, err := client.SayHelloServerStreamning(context.Background(), name)
// 	if err != nil {
// 		log.Fatalf("could not send names : %v", err)

// 	}
// 	for {
// 		mess, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Could not obtain the data stream", err)
// 		}
// 		log.Println(mess)

// 	}

// 	log.Println("Streaming Finished")
// }
