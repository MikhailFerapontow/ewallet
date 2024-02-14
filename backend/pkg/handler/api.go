package handler

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) newWallet(c *gin.Context) {
	var input models.Wallet

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Api.NewWallet()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) sendMoney(c *gin.Context) {

}

func (h *Handler) getHistory(c *gin.Context) {

}

func (h *Handler) getWallet(c *gin.Context) {

}
