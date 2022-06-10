package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	List(*gin.Context)
	Read(*gin.Context)
	Add(*gin.Context)
	Update(*gin.Context)
	Remove(*gin.Context)
}

type HandlerFunc interface{}
