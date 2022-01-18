package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/et-nik/otus-highload/internal/di"
	"github.com/et-nik/otus-highload/internal/domain"
	"github.com/gorilla/mux"
)

type UsersIDFriendsHandler struct {
	userRepository domain.UserRepository
}

func NewUsersIDFriendsHandler(c *di.Container) *UsersIDFriendsHandler {
	return &UsersIDFriendsHandler{userRepository: c.UserRepository()}
}

func (handler *UsersIDFriendsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte("invalid request"))
		return
	}

	user, err := handler.userRepository.FindByID(request.Context(), id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("failed to find user"))
		return
	}
	if user == nil {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte("not found"))
		return
	}

	ff := newFriendsFinder(handler.userRepository)
	friends, err := ff.findFriendsForUser(request.Context(), user.Friends)
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

type friendsFinder struct {
	userRepository domain.UserRepository
}

func newFriendsFinder(userRepository domain.UserRepository) *friendsFinder {
	return &friendsFinder{userRepository: userRepository}
}

func (friendsFinder friendsFinder) findFriendsForUser(ctx context.Context, ids []int) ([]*domain.User, error) {
	friends := make([]*domain.User, 0, len(ids))
	for i := range ids {
		friend, err := friendsFinder.userRepository.FindByID(ctx, ids[i])
		if err != nil {
			return nil, err
		}

		friends = append(friends, friend)
	}

	return friends, nil
}
