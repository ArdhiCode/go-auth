package controller

import (
	"github.com/ArdhiCode/go-auth/internal/api/service"
	"github.com/ArdhiCode/go-auth/internal/dto"
	"github.com/ArdhiCode/go-auth/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuth(authService service.AuthService) AuthController {
	return &authController{authService}
}

func (c *authController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid body request", err).Send(ctx)
		return
	}

	res, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailedWithCode(401, "login failed", err).Send(ctx)
		return
	}

	response.NewSuccess("login success", res).Send(ctx)
}

func (c *authController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailedWithCode(400, "invalid body request", err).Send(ctx)
		return
	}

	res, err := c.authService.Register(ctx, req)
	if err != nil {
		response.NewFailedWithCode(401, "register failed", err).Send(ctx)
		return
	}

	response.NewSuccess("register success", res).Send(ctx)
}
