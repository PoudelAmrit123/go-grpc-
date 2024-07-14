package main

import (
	// "context"
	pb "github.com/PoudelAmrit123/proto"
	"log"

	// pb "github.com/PoudelAmrit123/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)

	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	// client := pb.GreetService_SayHelloBidirectionalStreamingClient(conn)

	names := pb.NameList{
		Names: []string{"Amritp", "Mausam", "Swisil", "Krishna", "Sangam"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client, &names)
	// CallSayHelloClientStream(client, &names)
	callSayHelloDirectionalStream(client, &names)

}
