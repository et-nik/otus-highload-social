package api

import (
	"encoding/json"
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/et-nik/otus-highload/pkg/web"
	"github.com/et-nik/otus-highload/pkg/web/responder"
)

type FollowUserHandler struct {
	userRepository domain.UserRepository
}

func NewFollowUserHandler(c *di.Container) *FollowUserHandler {
	return &FollowUserHandler{userRepository: c.UserRepository()}
}

func (handler *FollowUserHandler) ServeHTTP(wr http.ResponseWriter, rq *http.Request) {
	var command struct {
		ID int `json:"id"`
	}

	decoder := json.NewDecoder(rq.Body)
	err := decoder.Decode(&command)
	if err != nil {
		responder.WriteError(wr, rq, web.NewError(err, http.StatusBadRequest, "invalid request"))
		return
	}

	session := sessionFromContext(rq.Context())
	user := session.User

	user.Friends = append(user.Friends, command.ID)

	err = handler.userRepository.Save(rq.Context(), user)
	if err != nil {
		responder.WriteError(wr, rq, web.NewServerInternalError(err, "failed to save user"))
		return
	}

	wr.WriteHeader(http.StatusOK)
}
