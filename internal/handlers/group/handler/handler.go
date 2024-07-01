package handler

import (
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/service/group/model"
	"github.com/Yakov-Varnaev/ft/internal/service/group/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	v1 := r.Group("v1")
	v1Group := v1.Group("groups")

	v1Group.GET("/", h.List)
	v1Group.POST("/", h.Post)
	v1Group.DELETE("/:id/", h.Delete)
	v1Group.PUT("/:id/", h.Put)
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(c *gin.Context) {
	var data model.GroupInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(&webErrors.BadRequest{Message: "Body cannot be empty."})
		return
	}
	group, err := h.service.Create(&data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Handler) List(c *gin.Context) {
	pg, err := pagination.NewFromRequest(c)
	if err != nil {
		c.Error(&webErrors.BadRequest{Message: err.Error()})
		return
	}
	groups, err := h.service.List(pg)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (h *Handler) Put(c *gin.Context) {
	var data model.GroupInfo
	id := c.Param("id")
	if id == "" {
		panic("id is required here.")
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err)
		return
	}
	group, err := h.service.Update(id, &data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
