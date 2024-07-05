package handler

import (
	"net/http"

	service "github.com/Yakov-Varnaev/ft/internal/service/category"
	"github.com/Yakov-Varnaev/ft/internal/service/category/model"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	v1 := r.Group("v1")
	v1Cat := v1.Group("categories")

	v1Cat.GET("/", h.List)
	v1Cat.POST("/", h.Post)
	v1Cat.DELETE("/:id/", h.Delete)
	v1Cat.PUT("/:id/", h.Put)
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(c *gin.Context) {
	var data model.CategoryInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(&webErrors.BadRequest{Message: "Body cannot be empty."})
		return
	}
	cat, err := h.service.Create(&data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, cat)
}

func (h *Handler) List(c *gin.Context) {
	pg, err := pagination.NewFromRequest(c)
	if err != nil {
		c.Error(&webErrors.BadRequest{Message: err.Error()})
		return
	}
	cats, err := h.service.List(pg)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, cats)
}

func (h *Handler) Put(c *gin.Context) {
	var data model.CategoryInfo
	id := c.Param("id")
	if id == "" {
		panic("id is required here.")
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err)
		return
	}
	category, err := h.service.Update(id, &data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
