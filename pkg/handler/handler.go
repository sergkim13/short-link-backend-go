package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sergkim13/short-link-backend-go/pkg/service"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	links := router.Group("/links")
	{
		links.POST("/", h.createLink)
		links.GET("/:short", h.getLink)
	}

	return router
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
