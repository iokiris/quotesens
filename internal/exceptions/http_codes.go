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

var errorTypeToHTTPCode = map[ErrorType]int{
	ErrTypeInvalidInput: http.StatusBadRequest,
	ErrTypeNotFound:     http.StatusNotFound,
	ErrTypeConflict:     http.StatusConflict,
	ErrTypeUnauthorized: http.StatusUnauthorized,
	ErrTypeForbidden:    http.StatusForbidden,
}

func MapErrorTypeToHTTPCode(t ErrorType) int {
	if code, ok := errorTypeToHTTPCode[t]; ok {
		return code
	}
	return http.StatusInternalServerError
}
