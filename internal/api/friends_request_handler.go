package api

import (
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
)

type FriendsRequestHandler struct {
	userRepository domain.UserRepository
}

func NewFriendsRequestHandler(c *di.Container) *FriendsRequestHandler {
	return &FriendsRequestHandler{userRepository: c.UserRepository()}
}

func (handler *FriendsRequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}
