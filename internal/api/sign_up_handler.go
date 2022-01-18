package api

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/matthewhartstonge/argon2"
)

type SignUpHandler struct {
	userRepository domain.UserRepository
	argon          *argon2.Config
}

func NewSignUpHandler(c *di.Container) *SignUpHandler {
	return &SignUpHandler{
		userRepository: c.UserRepository(),
		argon:          c.Argon(),
	}
}

func (handler *SignUpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var command struct {
		Email     string   `json:"email"`
		Password  string   `json:"password"`
		Name      string   `json:"name"`
		Surname   string   `json:"surname"`
		Age       int      `json:"age"`
		Sex       string   `json:"sex"`
		Interests []string `json:"interests"`
		City      string   `json:"city"`
	}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&command)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("invalid request"))
		return
	}

	existsUser, err := handler.userRepository.FindByEmail(request.Context(), command.Email)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed"))
		return
	}
	if existsUser != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = writer.Write([]byte("user is already exists"))
		return
	}

	passwordHash, err := handler.argon.HashEncoded([]byte(command.Password))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed"))
		return
	}

	user := domain.NewUser(
		command.Email,
		string(passwordHash),
		command.Name,
		command.Surname,
		command.Age,
		command.Sex,
		command.Interests,
		command.City,
	)

	b := make([]byte, 256)
	_, err = rand.Read(b)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed"))
		return
	}

	user.AuthTokenHash = fmt.Sprintf("%x", sha256.Sum256(b))

	err = handler.userRepository.Save(request.Context(), user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to save user"))
		return
	}

	r := struct {
		AuthToken string       `json:"token"`
		User      *domain.User `json:"user"`
	}{
		AuthToken: fmt.Sprintf("%d|%s", user.ID, user.AuthTokenHash),
		User:      user,
	}

	result, err := json.Marshal(r)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to marshal result"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(result)
}
