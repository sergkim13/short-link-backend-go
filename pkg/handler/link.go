package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergkim13/short-link-backend-go/pkg/model"
)

func (h *Handler) addLink(c *gin.Context) {
	var input model.LinkCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), err.Error())

		return
	}

	shortURL, err := h.services.Link.MakeShort(input.OriginalURL)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong on server side", err.Error())

		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"short_url": shortURL,
	})
}

func (h *Handler) getLink(c *gin.Context) {
	shortURL := c.Params.ByName("short")
	originalURL, err := h.services.Link.GetLink(shortURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
				newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("original url for %s not found", shortURL), err.Error())
				return
			}
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong on server side", err.Error())
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)
}
