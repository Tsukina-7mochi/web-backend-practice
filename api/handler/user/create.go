package user

import (
	"encoding/json"
	"log"
	"main/handler"
	. "main/usecase"
	"net/http"
	"regexp"
)

type CreateUserHandler struct {
	userUsecase UserUsecase
}

func NewCreateUserHandler(userUsecase UserUsecase) CreateUserHandler {
	return CreateUserHandler{userUsecase}
}

type createUserBody struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

var namePattern = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

func (h CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bodyStr := make([]byte, r.ContentLength)
	r.Body.Read(bodyStr)

	var body createUserBody
	if err := json.Unmarshal(bodyStr, &body); err != nil {
		handler.RespondError(w, `Request body must be JSON object contains "name" and "display_name".`, http.StatusBadRequest)
		return
	}

	if len(body.Name) < 4 {
		handler.RespondError(w, `"name" must be longer than 4 chacaters.`, http.StatusBadRequest)
		return
	}
	if !namePattern.MatchString(body.Name) {
		handler.RespondError(w, `"name" must be alphanumeric and underscore.`, http.StatusBadRequest)
		return
	}
	if len(body.DisplayName) == 0 {
		handler.RespondError(w, `"display_name" must not be empty.`, http.StatusBadRequest)
		return
	}

	if err := h.userUsecase.Create(body.Name, body.DisplayName); err != nil {
		log.Printf("Failed to create user. %v", err)
		handler.RespondError(w, "Failed to create user.", http.StatusInternalServerError)
		return
	}

	handler.RespondOK(w, http.StatusCreated)
}
