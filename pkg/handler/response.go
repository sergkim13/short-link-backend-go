package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errResp struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string, origErrMessage string) {
	logrus.Error(origErrMessage)
	c.AbortWithStatusJSON(statusCode, errResp{Message: message})
}