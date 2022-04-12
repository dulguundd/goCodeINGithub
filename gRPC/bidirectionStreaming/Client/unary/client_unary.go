package main

import (
	"context"
	"fmt"
	pb "goCodeINGithub/gRPC/bidirectionStreaming/API"
	"google.golang.org/grpc"
	"log"
	"time"
)

//const address = "192.168.1.102:50051"
const address = "localhost:50051"

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
	start := time.Now()
	count := 6
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserManagementClient(conn)

	runningcount := 0
	for runningcount < count {
		CreateNewUserUnary(client, &pb.NewUser{
			Name: "dulguun",
			Age:  22,
		})
		runningcount++
	}
	serviceLatencyLogger(start)
}

func serviceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	log.Println(logMessage)
}
