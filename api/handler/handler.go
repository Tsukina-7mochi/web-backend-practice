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

func DeleteUser(db *mydb.MyDB) http.HandlerFunc {
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

		err = db.DeleteUser(userID)
		if err != nil {
			log.Println("Failed to delete user: %v", err)
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
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

func PatchTodoStatus(db *mydb.MyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		todoIDstr, ok := vars["todoID"]
		if !ok {
			log.Println("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		todoID, err := strconv.Atoi(todoIDstr)
		if err != nil {
			log.Println("Invalid todo id: %v", err)
			http.Error(w, "Invalid todo id", http.StatusBadRequest)
		}

		query := r.URL.Query()
		if query == nil {
			log.Println("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		doneStr := query.Get("done")
		done, err := strconv.ParseBool(doneStr)
		if err != nil {
			log.Println("Invalid done value: %v", err)
			http.Error(w, "Invalid done value", http.StatusBadRequest)
			return
		}

		err = db.PatchTodo(todoID, done)
		if err != nil {
			log.Println("Failed to patch todo: %v", err)
			http.Error(w, "Failed to patch todo", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteTodo(db *mydb.MyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		todoIDstr, ok := vars["todoID"]
		if !ok {
			log.Println("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		todoID, err := strconv.Atoi(todoIDstr)
		if err != nil {
			log.Println("Invalid todo id: %v", err)
			http.Error(w, "Invalid todo id", http.StatusBadRequest)
		}

		err = db.BulkDeleteTodos([]int{todoID})
		if err != nil {
			log.Println("Failed to delete todo: %v", err)
			http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
