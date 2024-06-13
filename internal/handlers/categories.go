package handlers

import (
	"errors"
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/models"
	"github.com/Yakov-Varnaev/ft/internal/services"
	"github.com/Yakov-Varnaev/ft/internal/web"
	"github.com/gin-gonic/gin"
)

type Categories struct {
	service *services.Categories
}

func (h *Categories) Create(c *gin.Context) {
	data := &models.WriteCategory{}
	err := c.ShouldBindJSON(data)
	if err != nil {
		c.AbortWithError(
			http.StatusBadRequest,
			err,
		)
	}
	category, err := h.service.Create(data)
	if err != nil {
		if errors.Is(err, &web.ValidationError{}) {
			c.AbortWithError(
				http.StatusBadRequest,
				err,
			)
			return
		} else {
			c.AbortWithError(
				http.StatusInternalServerError,
				err,
			)
			return
		}
	}
	c.JSON(http.StatusOK, category)
}

func (h *Categories) List(c *gin.Context) {
	categories, err := h.service.List()
	if err != nil {
		switch {
		case errors.Is(err, &web.ValidationError{}):
			c.AbortWithError(
				http.StatusBadRequest,
				err,
			)
			return
		default:
			c.AbortWithError(
				http.StatusInternalServerError,
				err,
			)
			return
		}
	}
	c.JSON(http.StatusOK, categories)
}
