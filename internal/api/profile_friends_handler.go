package api

import (
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/et-nik/otus-highload/pkg/web"
	"github.com/et-nik/otus-highload/pkg/web/responder"
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
		responder.WriteError(
			writer,
			request,
			web.NewServerInternalError(err, "failed to find friends"),
		)
		return
	}

	responder.WriteJson(writer, request, friends)
}
