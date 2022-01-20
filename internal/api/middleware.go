package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/et-nik/otus-highload/pkg/web"
	"github.com/et-nik/otus-highload/pkg/web/responder"
)

var (
	errInvalidAuthHeader = errors.New("invalid auth header")
	errInvalidToken      = errors.New("invalid token in auth header")
)

func NewAuthMiddleware(
	userRepository domain.UserRepository,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authorization := request.Header.Get("Authorization")

		token, err := parseToken(authorization)
		if err != nil {
			responder.WriteError(
				writer,
				request,
				web.NewError(err, http.StatusUnauthorized, "unauthorized"),
			)
			return
		}

		var user *domain.User
		user, err = userRepository.FindByID(request.Context(), token.UserID)
		if err != nil || user == nil {
			writer.WriteHeader(http.StatusUnauthorized)
			_, _ = writer.Write([]byte("unauthorized"))
			return
		}

		if user.AuthTokenHash != token.AuthHash {
			writer.WriteHeader(http.StatusUnauthorized)
			_, _ = writer.Write([]byte("unauthorized"))
			return
		}

		session := &Session{
			User: user,
		}

		next.ServeHTTP(writer, request.WithContext(contextWithSession(request.Context(), session)))
	})
}

type token struct {
	UserID   int
	AuthHash string
}

func parseToken(authHeader string) (*token, error) {
	if authHeader == "" || len(authHeader) < 6 {
		return nil, errInvalidAuthHeader
	}

	if strings.ToLower(authHeader[:6]) != "bearer" {
		return nil, errInvalidAuthHeader
	}

	t := strings.SplitN(authHeader[7:], "|", 2)

	userID, err := strconv.Atoi(t[0])
	if err != nil {
		return nil, errInvalidToken
	}

	return &token{
		UserID:   userID,
		AuthHash: t[1],
	}, nil
}
