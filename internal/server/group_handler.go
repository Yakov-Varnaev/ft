package server

import (
	"fmt"
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func getUUIDFromParam(c *gin.Context, key ...string) (string, error) {
	paramKey := "id"
	if len(key) > 0 {
		paramKey = key[0]
	}
	id := c.Param(paramKey)
	if _, err := uuid.Parse(id); err != nil {
		return "", &webErrors.BadRequest{
			Message: fmt.Sprintf("invalid uuid: %v", err),
		}
	}
	return id, nil
}

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

func NewGroupHandler(service service.GroupService) *groupHandler {
	return &groupHandler{service}
}

func (h *groupHandler) Create(c *gin.Context) {
	var data model.GroupInfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.Error(err)
	}
	group, err := h.service.Create(&data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *groupHandler) List(c *gin.Context) {
	pg, err := pagination.NewFromContext(c)
	if err != nil {
		c.Error(err)
		return
	}
	groups, err := h.service.List(pg)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (h *groupHandler) Put(c *gin.Context) {
	id, err := getUUIDFromParam(c)
	if err != nil {
		c.Error(err)
		return
	}
	var data model.GroupInfo
	c.ShouldBindJSON(&data)
	updateGroup, err := h.service.Update(id, &data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, updateGroup)
}

func (h *groupHandler) Delete(c *gin.Context) {
	id, err := getUUIDFromParam(c)
	if err != nil {
		c.Error(err)
		return
	}
	err = h.service.Delete(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
