package handler

import (
	"log/slog"
	"net/http"
	"server/internal/logger"
	"server/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) newWallet(c *gin.Context) {
	const op = "handler.newWallet"
	log := h.log.With(slog.String("op", op))

	wallet, err := h.services.Api.NewWallet()
	if err != nil {
		log.Debug("Error in request", logger.Err(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) sendMoney(c *gin.Context) {
	const op = "handler.sendMoney"
	log := h.log.With(slog.String("op", op))

	var input models.SendMoneyInput

	if err := c.BindJSON(&input); err != nil {
		log.Debug("Error in request", logger.Err(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fromId := c.Param("id")

	if err := h.services.Api.SendMoney(fromId, input); err != nil {
		log.Debug("Error in request", logger.Err(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (h *Handler) getHistory(c *gin.Context) {
	const op = "handler.getHistory"
	log := h.log.With(slog.String("op", op))

	id := c.Param("id")

	transactions, err := h.services.Api.GetHistory(id)
	if err != nil {
		log.Debug("Error in request", logger.Err(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) getWallet(c *gin.Context) {
	const op = "handler.getWallet"
	log := h.log.With(slog.String("op", op))

	id := c.Param("id")

	data, err := h.services.Api.GetWallet(id)
	if err != nil {
		log.Debug("Error in request", logger.Err(err))
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, data)
}
