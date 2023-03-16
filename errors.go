package qbittorrent_api

import (
	"github.com/pkg/errors"
	"net/http"
)

var (
	//ErrMissingRequiredParameters = errors.New("MissingRequiredParameters")
	ErrInvalidRequest       = errors.New("InvalidRequest")
	ErrUnauthorized         = errors.New("Unauthorized")
	ErrForbidden            = errors.New("Forbidden")
	ErrNotFound             = errors.New("NotFound")
	ErrMethodNotAllowed     = errors.New("MethodNotAllowed")
	ErrConflict             = errors.New("Conflict")
	ErrUnsupportedMediaType = errors.New("UnsupportedMediaType")
	ErrInternalServerError  = errors.New("InternalServerError")
)

func handleResponsesErr(statusCode int) error {
	if statusCode < 400 {
		return nil
	}

	switch statusCode {
	case http.StatusBadRequest:
		return ErrInvalidRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusMethodNotAllowed:
		return ErrMethodNotAllowed
	case http.StatusConflict:
		return ErrConflict
	case http.StatusUnsupportedMediaType:
		return ErrUnsupportedMediaType

	}
	return ErrInternalServerError
}
