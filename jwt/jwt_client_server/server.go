package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	pb "grpc-signup/proto/signuppb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSignupServiceServer
}

func (s *server) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	fmt.Printf("📩 Received Signup Request:\n📧 Gmail: %s\n🔐 Password: %s\n", req.Gmail, req.Password)
	return &pb.SignupResponse{Message: "Signup successful for " + req.Gmail}, nil
}

func main() {
	// Load server certificate and key
	cert, err := tls.LoadX509KeyPair("../certs/server.crt", "../certs/server.key")
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}

	// Create certificate pool from CA
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../certs/server.crt")
	if err != nil {
		log.Fatalf("could not read ca certificate: %v", err)
	}
	certPool.AppendCertsFromPEM(ca)

	// Create TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
	})

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterSignupServiceServer(grpcServer, &server{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("🚀 gRPC Server listening on https://localhost:50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
