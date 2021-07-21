package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mayerkv/otus-ma-lesson-6/domain"
	"net/http"
)

func handleError(ctx *gin.Context, err error) {
	code := http.StatusInternalServerError
	msg := "Internal server error"

	switch err.(type) {
	case *domain.EntityNotExists:
		code = http.StatusNotFound
		msg = err.Error()
	case *domain.InvalidArgument, validator.ValidationErrors:
		code = http.StatusBadRequest
		msg = err.Error()
	}

	ctx.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": msg,
	})
}
