package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sergkim13/short-link-backend-go/pkg/model"
)

func (h *Handler) addLink(ctx *gin.Context) {
	var input model.LinkCreate

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error(), err.Error())

		return
	}

	shortURL, err := h.services.Link.MakeShort(input.OriginalURL)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "something went wrong on server side", err.Error())

		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"short_url": shortURL,
	})
}

func (h *Handler) getLink(ctx *gin.Context) {
	shortURL := ctx.Params.ByName("short")
	originalURL, err := h.services.Link.GetLink(shortURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
				newErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("original url for %s not found", shortURL), err.Error())

				return
			}

		newErrorResponse(ctx, http.StatusInternalServerError, "something went wrong on server side", err.Error())

		return
	}

	ctx.Redirect(http.StatusMovedPermanently, originalURL)
}
