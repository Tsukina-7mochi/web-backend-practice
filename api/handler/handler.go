package handler

import (
	"fmt"
	"log"
	"main/mydb"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}

func AddUser(db *mydb.MyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if query == nil {
			log.Println("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		name := query.Get("name")

		id, err := db.AddUser(name)
		if err != nil {
			log.Println("Failed to add user: %v", err)
			http.Error(w, "Failed to add user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{id:%d}", id)
	}
}

func AddTodo(db *mydb.MyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		userIDstr, ok := vars["userID"]
		if !ok {
			log.Println("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(userIDstr)
		if err != nil {
			log.Println("Invalid user id: %v", err)
			http.Error(w, "Invalid user id", http.StatusBadRequest)
		}

		query := r.URL.Query()
		if query == nil {
			log.Println("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		title := query.Get("title")

		id, err := db.AddTodo(userID, title)
		if err != nil {
			log.Println("Failed to add todo: %v", err)
			http.Error(w, "Failed to add todo", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{id:%d}", id)
	}
}

func ListTodo(db *mydb.MyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userIDstr, ok := vars["userID"]
		if !ok {
			log.Println("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(userIDstr)
		if err != nil {
			log.Println("Invalid user id: %v", err)
			http.Error(w, "Invalid user id", http.StatusBadRequest)
		}

		todos, err := db.ListByUser(userID)
		if err != nil {
			log.Println("Failed to list todos: %v", err)
			http.Error(w, "Failed to list todos", http.StatusInternalServerError)
			return
		}

		todoStrings := make([]string, 0, len(todos))
		for _, todo := range todos {
			todoStrings = append(todoStrings, todo.JSON())
		}

		fmt.Fprintf(w, "[%s]", strings.Join(todoStrings, ","))
	}
}
