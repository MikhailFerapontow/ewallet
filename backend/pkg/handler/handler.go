package handler

import (
	"server/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1/wallet")
	{
		api.POST("/", h.newWallet)
		api.POST("/:id/send", h.sendMoney)
		api.GET("/:id/history", h.getHistory)
		api.GET("/:id", h.getWallet)
	}

	return router
}
