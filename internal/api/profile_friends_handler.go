package api

import (
	"encoding/json"
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
)

type ProfileFriendsHandler struct {
	userRepository domain.UserRepository
}

func NewProfileFriendsHandler(c *di.Container) *ProfileFriendsHandler {
	return &ProfileFriendsHandler{userRepository: c.UserRepository()}
}

func (handler *ProfileFriendsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	session := sessionFromContext(request.Context())

	ff := newFriendsFinder(handler.userRepository)
	friends, err := ff.findFriendsForUser(request.Context(), session.User.Friends)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to find friends"))
		return
	}

	result, err := json.Marshal(friends)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to marshal friends"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(result)
}
