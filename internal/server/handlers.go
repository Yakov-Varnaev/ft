package server

import (
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/service"
	"github.com/gin-gonic/gin"
)

type groupHandler struct {
	service service.GroupService
}

func (h *groupHandler) List(c *gin.Context) {
	groups, err := h.service.List()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
	}
	c.JSON(http.StatusOK, groups)
}

func NewGroupHandler(service service.GroupService) *groupHandler {
	return &groupHandler{service}
}
