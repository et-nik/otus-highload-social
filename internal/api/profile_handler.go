package api

import (
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/et-nik/otus-highload/pkg/web/responder"
)

type ProfileHandler struct {
	userRepository domain.UserRepository
}

func NewProfileHandler(c *di.Container) *ProfileHandler {
	return &ProfileHandler{userRepository: c.UserRepository()}
}

func (handler *ProfileHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	session := sessionFromContext(request.Context())
	responder.WriteJson(writer, request, session.User)
}
