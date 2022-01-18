package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/gorilla/mux"
)

type UsersIDHandler struct {
	userRepository domain.UserRepository
}

func NewUsersIDHandler(c *di.Container) *UsersIDHandler {
	return &UsersIDHandler{userRepository: c.UserRepository()}
}

func (handler *UsersIDHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("invalid request"))
		return
	}

	user, err := handler.userRepository.FindByID(request.Context(), id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to find user"))
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte("not found"))
		return
	}

	result, err := json.Marshal(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to marshal user"))
		return
	}

	_, _ = writer.Write(result)
}
