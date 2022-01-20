package responder

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/et-nik/otus-highload/pkg/web"
)

type Error interface {
	error
	HTTPStatus() int
	Message() string
}

func errorHTTPStatus(err error) int {
	var webError Error
	if errors.As(err, &webError) {
		return webError.HTTPStatus()
	}

	return http.StatusInternalServerError
}

func errorMessage(err error) string {
	var webError Error
	if errors.As(err, &webError) {
		return webError.Message()
	}

	return "Internal server error"
}

func WriteError(writer http.ResponseWriter, _ *http.Request, err error) {
	log.Println(err)

	writer.WriteHeader(errorHTTPStatus(err))
	_, writeErr := writer.Write([]byte(errorMessage(err)))
	if writeErr != nil {
		log.Println(err)
	}
}

func WriteJSON(writer http.ResponseWriter, request *http.Request, data interface{}) {
	result, err := json.Marshal(data)
	if err != nil {
		WriteError(
			writer,
			request,
			web.NewError(err, http.StatusInternalServerError, "failed to marshal data"),
		)
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(result)
	if err != nil {
		log.Println(err)
	}
}
