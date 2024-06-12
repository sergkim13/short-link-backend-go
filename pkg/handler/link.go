package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergkim13/short-link-backend-go/pkg/model"
)

func (h * Handler) addLink(c *gin.Context) {
	var input model.LinkCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	shortURL, err := h.services.Link.MakeShort(input.OriginalURL)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"short_url": shortURL,
	})

}

func (h *Handler) getLink(c *gin.Context) {
	url := c.Params.ByName("short")
	resp := fmt.Sprintf("Short link: %s", url)
	c.JSON(http.StatusOK, resp)
}
