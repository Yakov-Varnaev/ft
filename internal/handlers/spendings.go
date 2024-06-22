package handlers

import (
	"fmt"
	"net/http"

	"github.com/Yakov-Varnaev/ft/internal/models"
	"github.com/Yakov-Varnaev/ft/internal/services"
	"github.com/Yakov-Varnaev/ft/internal/web"
	"github.com/gin-gonic/gin"
)

type Spendings struct {
	service *services.Spendigns
}

func (h *Spendings) Create(c *gin.Context) {
	var data models.WriteSpendings
	err := c.ShouldBindJSON(&data)
	fmt.Printf("%v\n", data)
	if err != nil {
		web.ResponseError(c, err)
		return
	}

	spending, err := h.service.Create(&data)
	if err != nil {
		web.ResponseError(c, err)
		return
	}
	fmt.Printf("%v\n", spending)

	c.JSON(http.StatusOK, spending)
}

func (h *Spendings) List(c *gin.Context) {
	spendings, err := h.service.List()
	if err != nil {
		web.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, spendings)
}

func (h *Spendings) Update(c *gin.Context) {
	id := c.Param("id")
	data, err := web.GetDataFromBody[models.WriteSpendings](c)
	if err != nil {
		web.ResponseError(c, err)
		return
	}
	updatedSpending, err := h.service.Update(id, data)
	if err != nil {
		web.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedSpending)
}

func (h *Spendings) Delete(c *gin.Context) {
	id := c.Param("id")
	deletedId, err := h.service.Delete(id)
	if err != nil {
		web.ResponseError(c, err)
		return
	}

	c.JSON(http.StatusOK, deletedId)
}
