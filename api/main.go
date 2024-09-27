package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/handler"
	todoHandler "main/handler/todo"
	userHandler "main/handler/user"
	. "main/repository"
	. "main/usecase"
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
	if addr == "" {
		addr = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost,
			dbPort,
			dbUser,
			dbPass,
			dbName,
		),
	)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	userRepository := NewUserRepository(db)
	todoRepository := NewTodoRepository(db)

	userUsecase := NewUserUsecase(*userRepository)
	todoUsecase := NewTodoUsecase(*userRepository, *todoRepository)

	createUserHandler := userHandler.NewCreateUserHandler(*userUsecase)
	getUserHandler := userHandler.NewGetUserHandler(*userUsecase, "userName")
	deleteUserHandler := userHandler.NewDeleteUserHandler(*userUsecase, "userName")
	createTodoHandler := todoHandler.NewCreateTodoHandler(*todoUsecase, "userName")
	listTodoByUserNameHandler := todoHandler.NewListByUserNameHandler(*todoUsecase, "userName")
	updateTodoDoneHandler := todoHandler.NewUpdateDoneHandler(*todoUsecase, "userName", "todoRef")
	deleteTodoHandler := todoHandler.NewDeleteTodoHandler(*todoUsecase, "userName", "todoRef")

	r := mux.NewRouter()

	r.HandleFunc("/", handler.NotFound)
	r.HandleFunc("/ping", handler.Ping)
	r.Handle("/users", createUserHandler).Methods("POST")
	r.Handle("/users/{userName}", getUserHandler).Methods("GET")
	r.Handle("/users/{userName}", deleteUserHandler).Methods("DELETE")

	r.Handle("/users/{userName}/todos", listTodoByUserNameHandler).Methods("GET")
	r.Handle("/users/{userName}/todos", createTodoHandler).Methods("POST")
	r.Handle("/users/{userName}/todos/{todoRef}", updateTodoDoneHandler).Methods("PATCH")
	r.Handle("/users/{userName}/todos/{todoRef}", deleteTodoHandler).Methods("DELETE")

	log.Printf("Server listening on %v:%v", addr, port)
	log.Fatal(http.ListenAndServe(addr+":"+port, printLog(r)))
}
