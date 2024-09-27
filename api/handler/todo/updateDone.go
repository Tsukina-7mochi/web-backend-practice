package todo

import (
	"encoding/json"
	"log"
	"main/handler"
	. "main/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type UpdateDoneHandler struct {
	todoUsecase TodoUsecase
	UserNameKey string
	TodoRefKey  string
}

func NewUpdateDoneHandler(todoUsecase TodoUsecase, userNameKey string, todoRefKey string) UpdateDoneHandler {
	return UpdateDoneHandler{
		todoUsecase: todoUsecase,
		UserNameKey: userNameKey,
		TodoRefKey:  todoRefKey,
	}
}

type updateTodoDoneBody struct {
	Done bool `json:"done"`
}

func (h UpdateDoneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	bodyStr := make([]byte, r.ContentLength)
	r.Body.Read(bodyStr)

	var body updateTodoDoneBody
	if err := json.Unmarshal(bodyStr, &body); err != nil {
		handler.RespondError(w, `Request body must be JSON object contains "done".`, http.StatusBadRequest)
		return
	}

	if err := h.todoUsecase.UpdateDone(userName, ref, body.Done); err != nil {
		if err == NoTodoError {
			handler.RespondError(w, "Todo not found.", http.StatusNotFound)
			return
		}

		log.Printf("Failed to update done. %v", err)
		handler.RespondError(w, "Failed to update done.", http.StatusInternalServerError)
		return
	}

	handler.RespondOK(w, http.StatusOK)
}
