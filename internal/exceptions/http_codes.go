package exceptions

import "net/http"

const (
	ErrTypeUnknown      ErrorType = iota
	ErrTypeInvalidInput           // 400
	ErrTypeNotFound               // 404
	ErrTypeConflict               // 409
	ErrTypeUnauthorized           // 401
	ErrTypeForbidden              // 403
)

func MapErrorTypeToHTTPCode(t ErrorType) int {
	switch t {
	case ErrTypeInvalidInput:
		return http.StatusBadRequest
	case ErrTypeNotFound:
		return http.StatusNotFound
	case ErrTypeConflict:
		return http.StatusConflict
	case ErrTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrTypeForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
