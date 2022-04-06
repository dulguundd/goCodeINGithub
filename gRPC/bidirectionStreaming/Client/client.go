package main

import (
	"context"
	pb "goCodeINGithub/gRPC/bidirectionStreaming/API"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const address = "192.168.1.102:50051"

func CreateNewUser(client pb.UserManagementClient) {
	NewUsers := []*pb.NewUser{
		{Name: "dulguunA", Age: 20},
		{Name: "dulguunB", Age: 21},
		{Name: "dulguunC", Age: 22},
		{Name: "dulguunD", Age: 23},
		{Name: "dulguunE", Age: 24},
		{Name: "dulguunF", Age: 25},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.CreateNewUser(ctx)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message User: %s, Age: %d, ID: %d", in.Name, in.Age, in.Id)
		}
	}()
	for _, NewUser := range NewUsers {
		if err := stream.Send(NewUser); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserManagementClient(conn)

	CreateNewUser(client)
}
