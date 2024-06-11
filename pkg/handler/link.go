package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h * Handler) createLink(c *gin.Context) {
	c.JSON(http.StatusCreated, "Created")

}

func (h *Handler) getLink(c *gin.Context) {
	url := c.Params.ByName("short")
	resp := fmt.Sprintf("Short link: %s", url)
	c.JSON(http.StatusOK, resp)
}
