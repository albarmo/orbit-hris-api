package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PasswordResetController interface {
	Create(ctx *gin.Context)
	GetByToken(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type passwordResetController struct {
	svc service.PasswordResetService
}

func NewPasswordResetController(s service.PasswordResetService) PasswordResetController {
	return &passwordResetController{svc: s}
}

func (c *passwordResetController) Create(ctx *gin.Context) {
	var req dto.PasswordResetCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	out, err := c.svc.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create password reset", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success create password reset", out)
	ctx.JSON(http.StatusCreated, res)
}

func (c *passwordResetController) GetByToken(ctx *gin.Context) {
	t := ctx.Param("token")
	token, err := uuid.Parse(t)
	if err != nil {
		res := utils.BuildResponseFailed("invalid token", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	out, err := c.svc.FindByToken(ctx.Request.Context(), token)
	if err != nil {
		res := utils.BuildResponseFailed("not found", err.Error(), nil)
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := utils.BuildResponseSuccess("success get password reset", out)
	ctx.JSON(http.StatusOK, res)
}

func (c *passwordResetController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		res := utils.BuildResponseFailed("invalid id", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	if err := c.svc.Delete(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed("failed delete", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success delete", nil)
	ctx.JSON(http.StatusOK, res)
}
