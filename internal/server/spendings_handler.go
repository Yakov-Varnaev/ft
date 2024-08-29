package server

import (
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/service"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/gin-gonic/gin"
)

type spendingsHandler struct {
	service service.SpendingsService
}

func NewSpendingsHandler(serv service.SpendingsService) *spendingsHandler {
	return &spendingsHandler{serv}
}

func (h *spendingsHandler) Create(c *gin.Context) {
	var data model.SpendingsInfo
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.Error(err)
		return
	}
	spendings, err := h.service.Create(&data)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, spendings)
}

func (h *spendingsHandler) List(c *gin.Context) {
	pg, err := pagination.NewFromContext(c)
	if err != nil {
		c.Error(err)
		return
	}
	spendings, err := h.service.List(pg)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, spendings)
}
