package api

import (
	"encoding/json"
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
)

type UsersHandler struct {
	userRepository domain.UserRepository
}

func NewUsersHandler(c *di.Container) *UsersHandler {
	return &UsersHandler{userRepository: c.UserRepository()}
}

func (handler *UsersHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	users, err := handler.userRepository.Find(request.Context())
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to find users"))
		return
	}

	result, err := json.Marshal(users)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to marshal users"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(result)
}
