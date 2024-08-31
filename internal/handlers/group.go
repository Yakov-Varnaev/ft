package handlers

import (
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/web/utils"
	"github.com/gin-gonic/gin"
)

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
func (h *GroupHandler) Delete(c *gin.Context) {
	id, err := utils.GetUUIDFromParam(c, "id")
	if err != nil {
		c.Error(err)
		return
	}
	if err = h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
