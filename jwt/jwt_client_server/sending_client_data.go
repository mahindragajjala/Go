package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pb "grpc-signup/proto/signuppb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// Load CA cert
	caCert, err := ioutil.ReadFile("../certs/server.crt")
	if err != nil {
		log.Fatalf("could not read ca certificate: %v", err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	// Create credentials
	creds := credentials.NewClientTLSFromCert(certPool, "localhost")

	// Dial server with TLS
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSignupServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.Signup(ctx, &pb.SignupRequest{
		Gmail:    "mahindra@example.com",
		Password: "secure123",
	})
	if err != nil {
		log.Fatalf("Signup failed: %v", err)
	}
	fmt.Printf("✅ Server Response: %s\n", resp.Message)
}
