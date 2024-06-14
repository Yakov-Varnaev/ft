package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type NotFound struct {
	Message string
}

func (e *NotFound) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return e.Message
}

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return e.Message
}

func GetHTTPException(err error) *HTTPError {
	errMsg := err.Error()
	var code int

	switch {
	case errors.Is(err, &ValidationError{}):
		code = http.StatusBadRequest
	case errors.Is(err, &NotFound{}):
		code = http.StatusNotFound
	case errors.Is(err, &InternalServerError{}):
		code = http.StatusInternalServerError
	default:
		code = http.StatusTeapot
	}

	return &HTTPError{
		Code:    code,
		Message: errMsg,
	}
}

func ResponseError(c *gin.Context, err error) {
	httpErr := GetHTTPException(err)
	c.AbortWithStatusJSON(
		httpErr.Code,
		gin.H{"detail": httpErr.Message},
	)
}
