package handler

import (
	"log/slog"
	"server/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	log      *slog.Logger
}

func NewHandler(services *service.Service, logger *slog.Logger) *Handler {
	return &Handler{
		services: services,
		log:      logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	api := router.Group("/api/v1/wallet")
	{
		api.POST("/", h.newWallet)
		api.POST("/:id/send", h.sendMoney)
		api.GET("/:id/history", h.getHistory)
		api.GET("/:id", h.getWallet)
	}

	return router
}
