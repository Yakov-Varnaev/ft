package handlers

import (
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/web/utils"
	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *categoryHandler {
	return &categoryHandler{service}
}

func (h *categoryHandler) Create(c *gin.Context) {
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

func (h *categoryHandler) List(c *gin.Context) {
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

func (h *categoryHandler) Update(c *gin.Context) {
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

func (h *categoryHandler) Delete(c *gin.Context) {
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
