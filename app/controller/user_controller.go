package controller

import (
	"github.com/gin-gonic/gin"
	"mayerkv/otus-ma-lesson-6/domain"
	"net/http"
)

type UserController struct {
	userService *domain.UserService
}

func (c *UserController) Handlers() []Handler {
	return []Handler{
		{"", http.MethodPost, c.CreateUser},
		{"/:id", http.MethodGet, c.GetUser},
		{"/:id", http.MethodPut, c.UpdateUser},
		{"/:id", http.MethodDelete, c.DeleteUser},
	}
}

func NewUserController(userService *domain.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var cmd domain.CreateUserCommand
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		handleError(ctx, err)
		return
	}

	user, err := c.userService.CreateUser(&cmd)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusCreated, user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	var cmd domain.GetUserCommand
	if err := ctx.ShouldBindUri(&cmd); err != nil {
		handleError(ctx, err)
		return
	}

	user, err := c.userService.GetUser(&cmd)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var cmd domain.UpdateUserCommand
	_ = ctx.ShouldBindUri(&cmd)
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		handleError(ctx, err)
		return
	}

	user, err := c.userService.UpdateUser(&cmd)
	if err != nil {
		handleError(ctx, err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	var cmd domain.DeleteUserCommand
	if err := ctx.ShouldBindUri(&cmd); err != nil {
		handleError(ctx, err)
		return
	}

	if err := c.userService.DeleteUser(&cmd); err != nil {
		handleError(ctx, err)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
