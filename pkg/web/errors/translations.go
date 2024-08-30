package errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var errMap = map[string]string{
	"required":     "Field is required",
	"unique-name":  "Must be unique",
	"group-exists": "Group with given ID does not exist",
}

func Translate(err error) gin.H {
	validationErrors := err.(validator.ValidationErrors)
	h := gin.H{}
	for _, err := range validationErrors {
		errText, ok := errMap[err.Tag()]
		if !ok {
			errText = "Invalid field"
		} else {
			errText = fmt.Sprintf(errText)
		}
		h[err.Field()] = errText
	}
	return h
}
