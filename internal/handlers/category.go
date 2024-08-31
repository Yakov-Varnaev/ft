package handlers

import (
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/web/utils"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var data model.CategoryInfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.Error(err)
		return
	}
	category, err := h.service.Create(&data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) List(c *gin.Context) {
	pg, err := pagination.NewFromContext(c)
	if err != nil {
		c.Error(err)
		return
	}
	categories, err := h.service.List(pg)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := utils.GetUUIDFromParam(c)
	if err != nil {
		c.Error(err)
		return
	}
	var data model.CategoryInfo
	c.ShouldBindJSON(&data)
	updateGroup, err := h.service.Update(id, &data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, updateGroup)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := utils.GetUUIDFromParam(c)
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
