package models

import "github.com/gin-gonic/gin"

func Result(g *gin.Context, status int, msg string, data interface{}) *gin.Context {
	g.JSON(status, gin.H{
		"Code":   status,
		"Message": msg,
		"Data": data,
	})
	return g
}