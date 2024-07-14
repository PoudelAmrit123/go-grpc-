package main

import (
	// "context"
	"context"
	"io"
	"log"
	"time"

	pb "github.com/PoudelAmrit123/proto"
)

func callSayHelloDirectionalStream(client pb.GreetServiceClient, name *pb.NameList) {
	log.Printf("Bidirectonal Streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names : %v", err)
	}

	waitch := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receving the file ")
			}

			log.Printf("Received the request %v", message)
		}
		close(waitch)

	}()

	for _, name := range name.Names {

		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the steam from the server side %v", err)
		}
		time.Sleep(2 * time.Second)

	}
	stream.CloseSend()
	<-waitch
	log.Printf("Bidirection stream finished")

}
