package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/master/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/master/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/master/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	MasterController interface {
		// Departments
		CreateDepartment(ctx *gin.Context)
		GetDepartments(ctx *gin.Context)
		GetDepartmentByID(ctx *gin.Context)
		UpdateDepartment(ctx *gin.Context)
		DeleteDepartment(ctx *gin.Context)

		// Locations
		CreateLocation(ctx *gin.Context)
		GetLocations(ctx *gin.Context)
		GetLocationByID(ctx *gin.Context)
		UpdateLocation(ctx *gin.Context)
		DeleteLocation(ctx *gin.Context)

		// Positions
		CreatePosition(ctx *gin.Context)
		GetPositions(ctx *gin.Context)
		GetPositionByID(ctx *gin.Context)
		UpdatePosition(ctx *gin.Context)
		DeletePosition(ctx *gin.Context)
	}

	masterController struct {
		masterService    service.MasterService
		masterValidation *validation.MasterValidation
	}
)

func NewMasterController(s service.MasterService) MasterController {
	masterValidation := validation.NewMasterValidation()
	return &masterController{
		masterService:    s,
		masterValidation: masterValidation,
	}
}

// Departments
func (c *masterController) CreateDepartment(ctx *gin.Context) {
	var req dto.DepartmentCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterValidation.ValidateDepartmentCreateRequest(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	deptModel := entities.Department{
		Name:        req.Name,
		Description: req.Description,
	}

	result, err := c.masterService.CreateDepartment(ctx.Request.Context(), nil, deptModel)
	if err != nil {
		res := utils.BuildResponseFailed("failed create department", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create department", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *masterController) GetDepartments(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)

	page, err := c.masterService.FindDepartments(ctx.Request.Context(), nil, &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get departments", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) GetDepartmentByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.masterService.GetDepartmentByID(ctx.Request.Context(), nil, id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get department", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) UpdateDepartment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.DepartmentUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterValidation.ValidateDepartmentUpdateRequest(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	deptModel := entities.Department{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	result, err := c.masterService.UpdateDepartment(ctx.Request.Context(), nil, deptModel)
	if err != nil {
		res := utils.BuildResponseFailed("failed update department", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success update department", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) DeleteDepartment(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterService.DeleteDepartment(ctx.Request.Context(), nil, id); err != nil {
		res := utils.BuildResponseFailed("failed delete department", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete department", nil)
	ctx.JSON(http.StatusOK, res)
}

// Locations
func (c *masterController) CreateLocation(ctx *gin.Context) {
	var req dto.LocationCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterValidation.ValidateLocationCreateRequest(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	locModel := entities.Location{
		Name:         req.Name,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		RadiusMeters: req.RadiusMeters,
	}
	result, err := c.masterService.CreateLocation(ctx.Request.Context(), nil, locModel)
	if err != nil {
		res := utils.BuildResponseFailed("failed create location", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create location", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *masterController) GetLocations(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)

	page, err := c.masterService.FindLocations(ctx.Request.Context(), nil, &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get locations", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) GetLocationByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.masterService.GetLocationByID(ctx.Request.Context(), nil, id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get location", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) UpdateLocation(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.LocationUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterValidation.ValidateLocationUpdateRequest(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	locModel := entities.Location{
		ID:           id,
		Name:         req.Name,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		RadiusMeters: req.RadiusMeters,
	}
	result, err := c.masterService.UpdateLocation(ctx.Request.Context(), nil, locModel)
	if err != nil {
		res := utils.BuildResponseFailed("failed update location", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success update location", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) DeleteLocation(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterService.DeleteLocation(ctx.Request.Context(), nil, id); err != nil {
		res := utils.BuildResponseFailed("failed delete location", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete location", nil)
	ctx.JSON(http.StatusOK, res)
}

// Positions
func (c *masterController) CreatePosition(ctx *gin.Context) {
	var req dto.PositionCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterValidation.ValidatePositionCreateRequest(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	deptID, _ := uuid.Parse(req.DepartmentID)
	posModel := entities.Position{
		DepartmentID: deptID,
		Name:         req.Name,
		Level:        req.Level,
	}
	result, err := c.masterService.CreatePosition(ctx.Request.Context(), nil, posModel)
	if err != nil {
		res := utils.BuildResponseFailed("failed create position", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create position", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *masterController) GetPositions(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)

	page, err := c.masterService.FindPositions(ctx.Request.Context(), nil, &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get positions", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) GetPositionByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.masterService.GetPositionByID(ctx.Request.Context(), nil, id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get position", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) UpdatePosition(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.PositionUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterValidation.ValidatePositionUpdateRequest(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	deptID, _ := uuid.Parse(req.DepartmentID)
	posModel := entities.Position{
		ID:           id,
		DepartmentID: deptID,
		Name:         req.Name,
		Level:        req.Level,
	}
	result, err := c.masterService.UpdatePosition(ctx.Request.Context(), nil, posModel)
	if err != nil {
		res := utils.BuildResponseFailed("failed update position", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success update position", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *masterController) DeletePosition(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.masterService.DeletePosition(ctx.Request.Context(), nil, id); err != nil {
		res := utils.BuildResponseFailed("failed delete position", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete position", nil)
	ctx.JSON(http.StatusOK, res)
}
