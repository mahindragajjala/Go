package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	// Setup HTTP/2 transport without TLS
	transport := &http2.Transport{
		AllowHTTP: true, // Allow HTTP/2 without TLS
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr) // Plaintext dial
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("🛰️ Protocol Used:", resp.Proto)
	fmt.Println("📨 Response:", string(body))
}
