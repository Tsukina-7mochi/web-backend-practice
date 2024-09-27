package user

import (
	"log"
	"main/handler"
	. "main/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteUserHandler struct {
	userUsecase UserUsecase
	UserNameKey string
}

func NewDeleteUserHandler(userUsecase UserUsecase, userNameKey string) DeleteUserHandler {
	return DeleteUserHandler{
		userUsecase: userUsecase,
		UserNameKey: userNameKey,
	}
}

func (h DeleteUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars[h.UserNameKey]
	if name == "" {
		handler.RespondError(w, `User name must not be empty.`, http.StatusBadRequest)
		return
	}

	if err := h.userUsecase.Delete(name); err != nil {
		if err == NoUserError {
			handler.RespondError(w, "Failed to get user.", http.StatusInternalServerError)
			return
		}

		log.Printf("Failed to delete user. %v", err)
		handler.RespondError(w, "Failed to delete user.", http.StatusInternalServerError)
		return
	}

	handler.RespondOK(w, http.StatusOK)
}
