package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	addr := os.Getenv("ip")
	port := os.Getenv("port")

	if addr == "" {
		addr = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/ping", ping)
	fmt.Printf("Server listening on %v:%v", addr, port)
	log.Fatal(http.ListenAndServe(addr+":"+port, nil))
}
