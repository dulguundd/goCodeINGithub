package main

import (
	"context"
	pb "goCodeINGithub/gRPC/bidirectionStreaming/API"
	"google.golang.org/grpc"
	"log"
	"time"
)

const address = "192.168.1.102:50051"

func CreateNewUserUnary(client pb.UserManagementClient, newUser *pb.NewUser) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := client.CreateNewUserUnary(ctx, newUser)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf(`User Detail:
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserManagementClient(conn)

	for i := 0; i < 6; i++ {
		CreateNewUserUnary(client, &pb.NewUser{
			Name: "dulguun",
			Age:  22,
		})
	}
}
