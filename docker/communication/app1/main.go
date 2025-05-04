package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from app1!")
	})
	fmt.Println("App1 running on port 8080")
	http.ListenAndServe(":8080", nil)
}
