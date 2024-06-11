package handler

import "github.com/gin-gonic/gin"

type Handler struct {

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
