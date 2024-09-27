package todo

import (
	"log"
	"main/handler"
	. "main/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteTodoHandler struct {
	todoUsecase TodoUsecase
	UserNameKey string
	TodoRefKey  string
}

func NewDeleteTodoHandler(todoUsecase TodoUsecase, userNameKey string, todoRefKey string) DeleteTodoHandler {
	return DeleteTodoHandler{
		todoUsecase: todoUsecase,
		UserNameKey: userNameKey,
		TodoRefKey:  todoRefKey,
	}
}

func (h DeleteTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userName := vars[h.UserNameKey]
	if userName == "" {
		handler.RespondError(w, `User name must not be empty.`, http.StatusBadRequest)
		return
	}

	ref := vars[h.TodoRefKey]
	if ref == "" {
		handler.RespondError(w, `Todo ref must not be empty.`, http.StatusBadRequest)
		return
	}

	if err := h.todoUsecase.Delete(userName, ref); err != nil {
		if err == NoTodoError {
			handler.RespondError(w, "Todo not found.", http.StatusNotFound)
			return
		}

		log.Printf("Failed to delete todo. %v", err)
		handler.RespondError(w, "Failed to delete todo.", http.StatusInternalServerError)
		return
	}

	handler.RespondOK(w, http.StatusOK)
}
