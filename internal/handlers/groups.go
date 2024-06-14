package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/models"
	"github.com/Yakov-Varnaev/ft/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Groups struct {
	service *services.Groups
}

func (h *Groups) Create(c *gin.Context) {
	group := &models.Group{}

	err := c.ShouldBindJSON(group)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	group, err = h.service.Create(group)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, group)
}

func (h *Groups) List(c *gin.Context) {
	groups, err := h.service.List()
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, groups)
}

func (h *Groups) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	updGroup := models.WriteGroup{}
	if err := c.ShouldBindJSON(&updGroup); err != nil {
		slog.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	group, err := h.service.Update(id, updGroup)
	if err != nil {
		slog.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Groups) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err = h.service.Delete(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
