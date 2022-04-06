package main

import (
	"context"
	pb "goCodeINGithub/gRPC/bidirectionStreaming/API"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer

	user []*pb.User
}

func (s *UserManagementServer) CreateNewUserUnary(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Recieved: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   user_id,
	}, nil
}

func (s *UserManagementServer) CreateNewUser(stream pb.UserManagement_CreateNewUserServer) error {
	var user_id int32
	var user pb.User
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		user_id = int32(rand.Intn(1000))

		// Note: this copy prevents blocking other clients while serving this one.
		// We don't need to do a deep copy, because elements in the slice are
		// insert-only and never modified.
		user.Id = user_id
		user.Age = in.GetAge()
		user.Name = in.GetName()

		if err := stream.Send(&user); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
