package handler

import (
	"net/http"

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

}

func (h *Handler) getHistory(c *gin.Context) {

}

func (h *Handler) getWallet(c *gin.Context) {
	id := c.Param("id")

	data, err := h.services.Api.GetWallet(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"wallet": data,
	})
}
