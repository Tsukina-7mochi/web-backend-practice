package todo

import (
	"encoding/json"
	"log"
	"main/handler"
	. "main/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateTodoHandler struct {
	todoUsecase TodoUsecase
	UserNameKey string
}

func NewCreateTodoHandler(todoUsecase TodoUsecase, userNameKey string) CreateTodoHandler {
	return CreateTodoHandler{
		todoUsecase: todoUsecase,
		UserNameKey: userNameKey,
	}
}

type createTodoBody struct {
	Title string `json:"title"`
}

func (h CreateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userName := vars[h.UserNameKey]
	if userName == "" {
		handler.RespondError(w, `User name must not be empty.`, http.StatusBadRequest)
		return
	}

	bodyStr := make([]byte, r.ContentLength)
	r.Body.Read(bodyStr)

	var body createTodoBody
	if err := json.Unmarshal(bodyStr, &body); err != nil {
		handler.RespondError(w, `Request body must be JSON object contains "title".`, http.StatusBadRequest)
		return
	}

	log.Printf("%s, %s", userName, body.Title)

	err := h.todoUsecase.Create(userName, body.Title)
	if err == NoUserError {
		handler.RespondError(w, `User not found`, http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Failed to add todo. %v", err)
		handler.RespondError(w, `Failed to add todo`, http.StatusInternalServerError)
		return
	}

	handler.RespondOK(w, http.StatusCreated)
}
