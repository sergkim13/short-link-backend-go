package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergkim13/short-link-backend-go/pkg/model"
)

func (h * Handler) createLink(c *gin.Context) {
	var input model.LinkCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Link.CreateLink(input.OriginalURL)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// return
	c.JSON(http.StatusCreated, map[string]interface{}{"id": id})

}

func (h *Handler) getLink(c *gin.Context) {
	url := c.Params.ByName("short")
	resp := fmt.Sprintf("Short link: %s", url)
	c.JSON(http.StatusOK, resp)
}
