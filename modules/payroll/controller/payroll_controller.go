package controller

import (
	"net/http"

	"fmt"

	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	PayrollController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	payrollController struct {
		payrollService    service.PayrollService
		payrollValidation *validation.PayrollValidation
		db                             *gorm.DB
	}
)

func NewPayrollController(injector *do.Injector, s service.PayrollService) PayrollController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	payrollValidation := validation.NewPayrollValidation()
	return &payrollController{
		payrollService:    s,
		payrollValidation: payrollValidation,
		db:                             db,
	}
}

func (c *payrollController) Create(ctx *gin.Context) {
	var req dto.PayrollCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.payrollService.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create payroll", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create payroll", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *payrollController) GetAll(ctx *gin.Context) {
	// simple pagination params
	offset := 0
	limit := 0
	// parse query params if provided
	if v := ctx.Query("limit"); v != "" {
		// ignore parse errors for brevity
		fmt.Sscanf(v, "%d", &limit)
	}
	if v := ctx.Query("offset"); v != "" {
		fmt.Sscanf(v, "%d", &offset)
	}

	list, err := c.payrollService.FindAll(ctx.Request.Context(), offset, limit)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success get payroll list", list)
	ctx.JSON(http.StatusOK, res)
}

func (c *payrollController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.payrollService.FindByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get payroll", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success get payroll", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *payrollController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.PayrollUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.payrollService.Update(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update payroll", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success update payroll", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *payrollController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.payrollService.Delete(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed("failed delete payroll", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete payroll", nil)
	ctx.JSON(http.StatusOK, res)
}
