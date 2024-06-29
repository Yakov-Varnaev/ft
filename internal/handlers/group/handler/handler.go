package handler

import (
	"fmt"
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/service/group/model"
	"github.com/Yakov-Varnaev/ft/internal/service/group/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	v1 := r.Group("v1")
	v1Group := v1.Group("groups")
	v1Group.GET("/", h.Get)
	v1Group.POST("/", h.Post)
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Post(c *gin.Context) {
	var data model.GroupInfo
	fmt.Println("here we go")
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	group, err := h.service.Create(&data)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Handler) Get(c *gin.Context) {
	pg, err := pagination.NewFromRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"detail": err.Error(),
			},
		)
		return
	}
	groups, err := h.service.List(pg)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"detail": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, groups)
}
