package api

import (
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/et-nik/otus-highload/pkg/web"
	"github.com/et-nik/otus-highload/pkg/web/responder"
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
		responder.WriteError(
			writer,
			request,
			web.NewServerInternalError(err, "failed to find users"),
		)
		return
	}

	responder.WriteJSON(writer, request, users)
}
