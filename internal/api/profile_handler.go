package api

import (
	"encoding/json"
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
)

type ProfileHandler struct {
	userRepository domain.UserRepository
}

func NewProfileHandler(c *di.Container) *ProfileHandler {
	return &ProfileHandler{userRepository: c.UserRepository()}
}

func (handler *ProfileHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	session := sessionFromContext(request.Context())

	result, err := json.Marshal(session.User)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to marshal user"))
		return
	}

	_, _ = writer.Write(result)
}
