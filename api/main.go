package main

import (
	"log"
	"main/handler"
	"main/mydb"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func printLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%v: %v", req.Method, req.URL)
		handler.ServeHTTP(w, req)
	})
}

func main() {
	addr := os.Getenv("IP")
	port := os.Getenv("PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := mydb.Open(mydb.DBInit{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPass,
		Name:     dbName,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}

	if addr == "" {
		addr = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handler.NotFound)
	r.HandleFunc("/ping", handler.Pong)
	r.HandleFunc("/users", handler.AddUser(db)).Methods("POST")
	r.HandleFunc("/users/{userID}/todos", handler.AddTodo(db)).Methods("POST")
	r.HandleFunc("/users/{userID}/todos", handler.ListTodo(db))

	log.Printf("Server listening on %v:%v", addr, port)
	log.Fatal(http.ListenAndServe(addr+":"+port, printLog(r)))
}
