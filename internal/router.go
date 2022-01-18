package internal

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/et-nik/otus-highload/internal/api"
	"github.com/et-nik/otus-highload/internal/di"
	"github.com/gorilla/mux"
)

type route struct {
	Method           string
	Path             string
	Handler          http.Handler
	AllowGuestAccess bool
}

func routes(container *di.Container) []route {
	return []route{
		{
			Method:           http.MethodPost,
			Path:             "/sign-up",
			Handler:          api.NewSignUpHandler(container),
			AllowGuestAccess: true,
		},
		{
			Method:           http.MethodPost,
			Path:             "/sign-in",
			Handler:          api.NewSignInHandler(container),
			AllowGuestAccess: true,
		},
		{
			Method:  http.MethodGet,
			Path:    "/profile",
			Handler: api.NewProfileHandler(container),
		},
		{
			Method:  http.MethodGet,
			Path:    "/profile/friends",
			Handler: api.NewProfileFriendsHandler(container),
		},
		{
			Method:  http.MethodPut,
			Path:    "/profile/friends",
			Handler: api.NewFollowUserHandler(container),
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: api.NewUsersHandler(container),
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: api.NewUsersIDHandler(container),
		},
		{
			Method:  http.MethodGet,
			Path:    "/profile/friends",
			Handler: api.NewUsersIDFriendsHandler(container),
		},
	}
}

func createRouter(container *di.Container) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	routeList := routes(container)

	for _, r := range routeList {
		handler := r.Handler

		if !r.AllowGuestAccess {
			handler = api.NewAuthMiddleware(container.UserRepository(), handler)
		}

		router.Handle(r.Path, handler).Methods(r.Method)
	}

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(struct {
			ApplicationVersion string `json:"application_version"`
			UserAgent          string `json:"user_agent"`
			Timestamp          string `json:"timestamp"`
		}{
			UserAgent: request.UserAgent(),
			Timestamp: time.Now().Format(time.RFC3339),
		})

		_, _ = writer.Write(response)
	})

	return router
}
