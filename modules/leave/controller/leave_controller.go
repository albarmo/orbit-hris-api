package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/leave/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	LeaveController interface {
		GetAll(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	leaveController struct {
		leaveService    service.LeaveService
		leaveValidation *validation.LeaveValidation
		db              *gorm.DB
	}
)

func NewLeaveController(injector *do.Injector, s service.LeaveService) LeaveController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	leaveValidation := validation.NewLeaveValidation()
	return &leaveController{
		leaveService:    s,
		leaveValidation: leaveValidation,
		db:              db,
	}
}

func (c *leaveController) GetAll(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)

	page, err := c.leaveService.FindAll(ctx.Request.Context(), &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get leaves", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

func (c *leaveController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.leaveService.GetByID(id)
	if err != nil {
		res := utils.BuildResponseFailed("not found", err.Error(), nil)
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *leaveController) Create(ctx *gin.Context) {
	var req dto.LeaveCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.leaveValidation.ValidateCreate(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.leaveService.Create(req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create leave", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create leave", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *leaveController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.LeaveUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.leaveValidation.ValidateUpdate(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.leaveService.Update(id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update leave", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("success update leave", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *leaveController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.leaveService.Delete(id); err != nil {
		res := utils.BuildResponseFailed("failed delete leave", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete leave", nil)
	ctx.JSON(http.StatusOK, res)
}
