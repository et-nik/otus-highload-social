package api

import (
	"encoding/json"
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
)

type FollowUserHandler struct {
	userRepository domain.UserRepository
}

func NewFollowUserHandler(c *di.Container) *FollowUserHandler {
	return &FollowUserHandler{userRepository: c.UserRepository()}
}

func (handler *FollowUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var command struct {
		ID int `json:"id"`
	}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&command)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("invalid request"))
		return
	}

	session := sessionFromContext(request.Context())
	user := session.User

	user.Friends = append(user.Friends, command.ID)

	err = handler.userRepository.Save(request.Context(), user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("failed to save user"))
		return
	}

	writer.WriteHeader(http.StatusOK)
}
