package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:message`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	// TODO: change logrus to slog or vice versa
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{Message: message})
}
