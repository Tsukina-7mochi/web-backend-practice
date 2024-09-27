package user

import (
	"log"
	"main/handler"
	. "main/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type GetUserHandler struct {
	userUsecase UserUsecase
	UserNameKey string
}

func NewGetUserHandler(userUsecase UserUsecase, userNameKey string) GetUserHandler {
	return GetUserHandler{
		userUsecase: userUsecase,
		UserNameKey: userNameKey,
	}
}

func (h GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars[h.UserNameKey]
	if name == "" {
		handler.RespondError(w, `User name must not be empty.`, http.StatusBadRequest)
		return
	}

	user, err := h.userUsecase.Get(name)
	if err == NoUserError {
		handler.RespondError(w, "User not found.", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Failed to get user. %v", err)
		handler.RespondError(w, "Failed to get user.", http.StatusInternalServerError)
		return
	}

	handler.RespondJSON(w, user, http.StatusOK)
}
