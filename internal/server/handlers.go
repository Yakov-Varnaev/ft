package server

import (
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/service"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func errorHandler(c *gin.Context) {
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

type groupHandler struct {
	service service.GroupService
}

func (h *groupHandler) List(c *gin.Context) {
	groups, err := h.service.List()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (h *groupHandler) Create(c *gin.Context) {
	var data model.GroupInfo
	c.ShouldBindJSON(&data)
	group, err := h.service.Create(&data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, group)
}

func NewGroupHandler(service service.GroupService) *groupHandler {
	return &groupHandler{service}
}
