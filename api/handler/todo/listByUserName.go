package todo

import (
	"log"
	"main/handler"
	. "main/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type ListByUserNameHandler struct {
	todoUsecase TodoUsecase
	UserNameKey string
}

func NewListByUserNameHandler(todoUsecase TodoUsecase, userNameKey string) ListByUserNameHandler {
	return ListByUserNameHandler{
		todoUsecase: todoUsecase,
		UserNameKey: userNameKey,
	}
}

func (h ListByUserNameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userName := vars[h.UserNameKey]
	if userName == "" {
		handler.RespondError(w, "User name must not be empty.", http.StatusBadRequest)
		return
	}

	todos, err := h.todoUsecase.ListByUser(userName)
	if err == NoUserError {
		handler.RespondError(w, "User not found.", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Failed to get user. %v", err)
		handler.RespondError(w, "Failed to get user.", http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(w, todos, http.StatusOK)
}
