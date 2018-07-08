package main

import (
	"log"
	"net"

	pb "github.com/ChristianVestre/nodeservice/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) GiveFirstName(ctx context.Context, in *pb.FirstNameRequest) (*pb.FirstNameReply, error) {
	return &pb.FirstNameReply{FirstNameReply: "Hello " + in.FirstName}, nil
}

func (s *server) GiveLastName(ctx context.Context, in *pb.LastNameRequest) (*pb.LastNameReply, error) {
	return &pb.LastNameReply{LastNameReply: "Hello " + in.LastName}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
