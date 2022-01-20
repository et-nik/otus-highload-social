package api

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/et-nik/otus-highload/pkg/web"
	"github.com/et-nik/otus-highload/pkg/web/responder"
	"github.com/matthewhartstonge/argon2"
)

type SignInHandler struct {
	userRepository domain.UserRepository
}

func NewSignInHandler(c *di.Container) *SignInHandler {
	return &SignInHandler{
		userRepository: c.UserRepository(),
	}
}

func (handler *SignInHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var command struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&command)
	if err != nil {
		responder.WriteError(
			writer,
			request,
			web.NewError(err, http.StatusBadRequest, "invalid request"),
		)
		return
	}

	user, err := handler.userRepository.FindByEmail(request.Context(), command.Email)
	if err != nil {
		responder.WriteError(
			writer,
			request,
			web.NewServerInternalError(err, "failed to find user"),
		)
		return
	}
	if user == nil {
		responder.WriteError(
			writer,
			request,
			web.NewError(err, http.StatusUnprocessableEntity, "failed to sign in"),
		)
		return
	}

	rawArgon, err := argon2.Decode([]byte(user.PasswordHash))
	if err != nil {
		responder.WriteError(
			writer,
			request,
			web.NewServerInternalError(err, "failed"),
		)
		return
	}

	ok, err := rawArgon.Verify([]byte(command.Password))
	if err != nil {
		responder.WriteError(
			writer,
			request,
			web.NewServerInternalError(err, "failed"),
		)
		return
	}

	if !ok {
		responder.WriteError(
			writer,
			request,
			web.NewError(err, http.StatusUnauthorized, "invalid credentials"),
		)
		return
	}

	b := make([]byte, 256)
	_, err = rand.Read(b)
	if err != nil {
		responder.WriteError(
			writer,
			request,
			web.NewServerInternalError(err, "failed"),
		)
		return
	}

	user.AuthTokenHash = fmt.Sprintf("%x", sha256.Sum256(b))

	err = handler.userRepository.Save(request.Context(), user)
	if err != nil {
		responder.WriteError(
			writer,
			request,
			web.NewServerInternalError(err, "failed"),
		)
		return
	}

	r := struct {
		AuthToken string       `json:"token"`
		User      *domain.User `json:"user"`
	}{
		AuthToken: fmt.Sprintf("%d|%s", user.ID, user.AuthTokenHash),
		User:      user,
	}

	responder.WriteJSON(writer, request, r)
}
