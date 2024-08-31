package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func printLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%v: %v", req.Method, req.URL)
		handler.ServeHTTP(w, req)
	})
}

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
	log.Printf("Server listening on %v:%v", addr, port)
	log.Fatal(http.ListenAndServe(addr+":"+port, printLog(http.DefaultServeMux)))
}
