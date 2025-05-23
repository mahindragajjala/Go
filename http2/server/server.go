package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("📡 Server received:", r.Proto)
	w.Write([]byte("Hello from HTTP/2 Cleartext Server!"))
}

func main() {
	h := http.HandlerFunc(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(h, &http2.Server{}), // <-- Enable h2c
	}

	fmt.Println("🚀 h2c server started on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
