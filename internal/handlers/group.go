package handlers

import (
	"fmt"
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/gin-gonic/gin"
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

type GroupHandler struct {
	service service.GroupService
}

func NewGroupHandler(service service.GroupService) *GroupHandler {
	return &GroupHandler{service}
}

func (h *GroupHandler) Create(c *gin.Context) {
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

func (h *GroupHandler) List(c *gin.Context) {
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
