package server

import (
	"net/http"

	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		switch e := err.Err.(type) {
		case *webErrors.InternalServerError:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		case *webErrors.NotFoundError:
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"message": e.Error()},
			)
			return
		case validator.ValidationErrors:
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				webErrors.Translate(e),
			)
			return
		case *webErrors.BadRequest:
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"message": e.Error()},
			)
			return
		default:
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}
