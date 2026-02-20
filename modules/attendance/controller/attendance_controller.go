package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/attendance/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	AttendanceController interface {
		GetAll(ctx *gin.Context)
		GetByID(ctx *gin.Context)
		GetByEmployeeID(ctx *gin.Context)
		CheckIn(ctx *gin.Context)
		CheckOut(ctx *gin.Context)
		Update(ctx *gin.Context)
		Approve(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	attendanceController struct {
		service    service.AttendanceService
		validation *validation.AttendanceValidation
		db         *gorm.DB
	}
)

func NewAttendanceController(injector *do.Injector, s service.AttendanceService) AttendanceController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	return &attendanceController{
		service:    s,
		validation: validation.NewAttendanceValidation(),
		db:         db,
	}
}

// GetAll godoc
// @Summary Get all attendances
// @Description Get all attendances
// @Tags attendances
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} utils.Response
// @Router /attendances [get]
func (c *attendanceController) GetAll(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)

	page, err := c.service.FindAll(ctx.Request.Context(), &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get attendances", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

// GetByID godoc
// @Summary Get attendance by ID
// @Description Get attendance by ID
// @Tags attendances
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Success 200 {object} utils.Response
// @Router /attendances/{id} [get]
func (c *attendanceController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.service.GetByID(id)
	if err != nil {
		res := utils.BuildResponseFailed("not found", err.Error(), nil)
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *attendanceController) GetByEmployeeID(ctx *gin.Context) {
	employeeID := ctx.Param("employee_id")

	var filter = pagination.Filter{}
	filter.Bind(ctx)

	page, err := c.service.FindByEmployeeID(ctx.Request.Context(), employeeID, &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get attendances by employee", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

// CheckIn godoc
// @Summary Check in
// @Description Check in
// @Tags attendances
// @Accept json
// @Produce json
// @Param body body dto.CheckInDTO true "CheckIn DTO"
// @Success 201 {object} utils.Response
// @Router /attendances/check-in [post]
func (c *attendanceController) CheckIn(ctx *gin.Context) {
	var req dto.CheckInDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.validation.CheckIn(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.service.CheckIn(req)
	if err != nil {
		res := utils.BuildResponseFailed("failed check-in", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("check-in successful", result)
	ctx.JSON(http.StatusCreated, res)
}

// CheckOut godoc
// @Summary Check out
// @Description Check out
// @Tags attendances
// @Accept json
// @Produce json
// @Param body body dto.CheckOutDTO true "CheckOut DTO"
// @Success 200 {object} utils.Response
// @Router /attendances/check-out [put]
func (c *attendanceController) CheckOut(ctx *gin.Context) {
	var req dto.CheckOutDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.validation.CheckOut(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.service.CheckOut(req)
	if err != nil {
		res := utils.BuildResponseFailed("failed check-out", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("check-out successful", result)
	ctx.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary Update attendance
// @Description Update attendance
// @Tags attendances
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Param body body dto.UpdateAttendanceDTO true "UpdateAttendance DTO"
// @Success 200 {object} utils.Response
// @Router /attendances/{id} [put]
func (c *attendanceController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.UpdateAttendanceDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.validation.UpdateAttendance(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.service.Update(id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update attendance", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("update successful", result)
	ctx.JSON(http.StatusOK, res)
}

// Approve godoc
// @Summary Approve attendance
// @Description Approve attendance
// @Tags attendances
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Param body body dto.ApproveAttendanceDTO true "ApproveAttendance DTO"
// @Success 200 {object} utils.Response
// @Router /attendances/{id}/approve [put]
func (c *attendanceController) Approve(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.ApproveAttendanceDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.validation.ApproveAttendance(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.service.Approve(id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed approval", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("approval successful", result)
	ctx.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Delete attendance
// @Description Delete attendance
// @Tags attendances
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Success 200 {object} utils.Response
// @Router /attendances/{id} [delete]
func (c *attendanceController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.service.Delete(id)
	if err != nil {
		res := utils.BuildResponseFailed("failed delete attendance", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.BuildResponseSuccess("delete successful", nil)
	ctx.JSON(http.StatusOK, res)
}
