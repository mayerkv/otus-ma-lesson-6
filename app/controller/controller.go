package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	Handlers() []Handler
}

type Handler struct {
	Path   string
	Method string
	gin.HandlerFunc
}
