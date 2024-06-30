package handler

import (
	"fmt"
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/service/group/model"
	"github.com/Yakov-Varnaev/ft/internal/service/group/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

var errMap = map[string]string{
	"required": "Field %s is required",
}

var tagPrefixMap = map[string]string{
	"required": "Required",
	"email":    "InvalidEmail",
	"min":      "ShouldMin",
	"max":      "ShouldMax",
	"len":      "ShouldLen",
	"eq":       "ShouldEq",
	"gt":       "ShouldGt",
	"gte":      "ShouldGte",
	"lt":       "ShouldLt",
	"lte":      "ShouldLte",
}

func i18n(msgTemplate string, params ...interface{}) string {
	return fmt.Sprintf(msgTemplate, params...)
}

func composeMsg(e validator.FieldError) string {
	if prefix, ok := tagPrefixMap[e.Tag()]; ok {
		return fmt.Sprintf("%s%s", prefix, e.Field())
	}
	return ""
}

func translateError(err error) string {
	var errTxt string
	validationErrors := err.(validator.ValidationErrors)
	for _, e := range validationErrors {
		errTxt = i18n(composeMsg(e), e.Param())
		break
	}
	return errTxt
}

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
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	group, err := h.service.Create(&data)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Handler) List(c *gin.Context) {
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

func (h *Handler) Put(c *gin.Context) {
	var data model.GroupInfo
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"detail": err.Error()},
		)
		return
	}
	group, err := h.service.Update(id, &data)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.Status(http.StatusNoContent)
}
