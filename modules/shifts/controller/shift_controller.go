package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/shifts/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/shifts/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShiftController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)

	CreateEmployeeShift(ctx *gin.Context)
	GetEmployeeShifts(ctx *gin.Context)
	DeleteEmployeeShift(ctx *gin.Context)
}

type shiftController struct {
	svc service.ShiftService
}

func NewShiftController(s service.ShiftService) ShiftController {
	return &shiftController{svc: s}
}

func (c *shiftController) Create(ctx *gin.Context) {
	var req dto.ShiftCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	out, err := c.svc.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create shift", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success create", out)
	ctx.JSON(http.StatusCreated, res)
}

func (c *shiftController) GetAll(ctx *gin.Context) {
	list, err := c.svc.FindAll(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("failed get list", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success get list", list)
	ctx.JSON(http.StatusOK, res)
}

func (c *shiftController) GetByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	out, err := c.svc.FindByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success get", out)
	ctx.JSON(http.StatusOK, res)
}

func (c *shiftController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var req dto.ShiftUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	out, err := c.svc.Update(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success update", out)
	ctx.JSON(http.StatusOK, res)
}

func (c *shiftController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
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

func (c *shiftController) CreateEmployeeShift(ctx *gin.Context) {
	var req dto.EmployeeShiftCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	out, err := c.svc.CreateEmployeeShift(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create employee shift", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success create", out)
	ctx.JSON(http.StatusCreated, res)
}

func (c *shiftController) GetEmployeeShifts(ctx *gin.Context) {
	eid, err := uuid.Parse(ctx.Param("employee_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	out, err := c.svc.FindEmployeeShifts(ctx.Request.Context(), eid)
	if err != nil {
		res := utils.BuildResponseFailed("failed get", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success get", out)
	ctx.JSON(http.StatusOK, res)
}

func (c *shiftController) DeleteEmployeeShift(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	if err := c.svc.DeleteEmployeeShift(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed("failed delete", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success delete", nil)
	ctx.JSON(http.StatusOK, res)
}
