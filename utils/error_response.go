package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest(c *gin.Context, msg string) *gin.Context {
	status:= http.StatusBadRequest
	c.JSON(status,gin.H{
		"Code" : status,
		"Message": msg,
	})
	return c
}

func InternalServerError(c *gin.Context, msg string) *gin.Context {
	status:= http.StatusInternalServerError
	c.JSON(status,gin.H{
		"Code" : status,
		"Message": msg,
	})
	return c
}

func NotFound(c *gin.Context, msg string) *gin.Context {
	status:= http.StatusNotFound
	c.JSON(status,gin.H{
		"Code" : status,
		"Message": msg,
	})
	return c
}

