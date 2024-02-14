package handler

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) newWallet(c *gin.Context) {
	id, err := h.services.Api.NewWallet()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) sendMoney(c *gin.Context) {
	var input models.SendMoneyInput

	if err := c.BindJSON(&input); err != nil {
		op := " binding" //УБРАТЬ
		newErrorResponse(c, http.StatusBadRequest, err.Error()+op)
		return
	}

	fromId := c.Param("id")

	if err := h.services.Api.SendMoney(fromId, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) getHistory(c *gin.Context) {
	id := c.Param("id")

	transactions, err := h.services.Api.GetHistory(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) getWallet(c *gin.Context) {
	id := c.Param("id")

	data, err := h.services.Api.GetWallet(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
