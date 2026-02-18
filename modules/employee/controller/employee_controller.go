package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmployeeController interface {
	Create(ctx *gin.Context)
	GetAllEmployee(ctx *gin.Context)
	GetEmployeeByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	// Child table handlers
	CreatePersonalInfo(ctx *gin.Context)
	GetPersonalInfo(ctx *gin.Context)
	UpdatePersonalInfo(ctx *gin.Context)
	DeletePersonalInfo(ctx *gin.Context)

	CreateAddress(ctx *gin.Context)
	GetAddresses(ctx *gin.Context)
	GetAddressByID(ctx *gin.Context)
	UpdateAddress(ctx *gin.Context)
	DeleteAddress(ctx *gin.Context)

	CreateLegalInfo(ctx *gin.Context)
	GetLegalInfo(ctx *gin.Context)
	UpdateLegalInfo(ctx *gin.Context)
	DeleteLegalInfo(ctx *gin.Context)

	CreatePayrollProfile(ctx *gin.Context)
	GetPayrollProfile(ctx *gin.Context)
	UpdatePayrollProfile(ctx *gin.Context)
	DeletePayrollProfile(ctx *gin.Context)
}

type employeeController struct {
	employeeService    service.EmployeeService
	employeeValidation *validation.EmployeeValidation
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &employeeController{
		employeeService:    employeeService,
		employeeValidation: validation.NewEmployeeValidation(),
	}
}

func (c *employeeController) Create(ctx *gin.Context) {
	var req dto.EmployeeCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.employeeValidation.ValidateEmployeeCreateRequest(req); err != nil {
		res := utils.BuildResponseFailed("Validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.Create(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_EMPLOYEE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_EMPLOYEE, result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *employeeController) GetAllEmployee(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)

	employees, err := c.employeeService.FindAll(ctx.Request.Context(), &filter)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_EMPLOYEE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_EMPLOYEE, employees)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) GetEmployeeByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.FindByID(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_EMPLOYEE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_EMPLOYEE, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) Update(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeeUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.employeeValidation.ValidateEmployeeUpdateRequest(req); err != nil {
		res := utils.BuildResponseFailed("Validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.Update(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_EMPLOYEE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_EMPLOYEE, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.employeeService.Delete(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_EMPLOYEE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_EMPLOYEE, nil)
	ctx.JSON(http.StatusOK, res)
}

// Personal Info handlers
func (c *employeeController) CreatePersonalInfo(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeePersonalInfoCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.CreatePersonalInfo(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create personal info", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create personal info", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *employeeController) GetPersonalInfo(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.GetPersonalInfo(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get personal info", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success get personal info", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) UpdatePersonalInfo(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeePersonalInfoUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.UpdatePersonalInfo(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update personal info", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success update personal info", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) DeletePersonalInfo(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.employeeService.DeletePersonalInfo(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed("failed delete personal info", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete personal info", nil)
	ctx.JSON(http.StatusOK, res)
}

// Addresses handlers
func (c *employeeController) CreateAddress(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeeAddressCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.CreateAddress(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create address", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create address", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *employeeController) GetAddresses(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.GetAddresses(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get addresses", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success get addresses", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) GetAddressByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("address_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.GetAddress(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get address", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success get address", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) UpdateAddress(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("address_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeeAddressUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// ensure id set
	if req.ID == nil {
		req.ID = &id
	}

	result, err := c.employeeService.UpdateAddress(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update address", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success update address", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) DeleteAddress(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("address_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.employeeService.DeleteAddress(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed("failed delete address", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete address", nil)
	ctx.JSON(http.StatusOK, res)
}

// Legal info handlers
func (c *employeeController) CreateLegalInfo(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeeLegalInfoCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.CreateLegalInfo(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create legal info", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create legal info", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *employeeController) GetLegalInfo(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.GetLegalInfo(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get legal info", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success get legal info", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) UpdateLegalInfo(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeeLegalInfoUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.UpdateLegalInfo(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update legal info", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success update legal info", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) DeleteLegalInfo(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.employeeService.DeleteLegalInfo(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed("failed delete legal info", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete legal info", nil)
	ctx.JSON(http.StatusOK, res)
}

// Payroll handlers
func (c *employeeController) CreatePayrollProfile(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeePayrollProfileCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.CreatePayrollProfile(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed create payroll profile", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success create payroll profile", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *employeeController) GetPayrollProfile(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.GetPayrollProfile(ctx.Request.Context(), id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get payroll profile", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success get payroll profile", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) UpdatePayrollProfile(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var req dto.EmployeePayrollProfileUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.employeeService.UpdatePayrollProfile(ctx.Request.Context(), id, req)
	if err != nil {
		res := utils.BuildResponseFailed("failed update payroll profile", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success update payroll profile", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *employeeController) DeletePayrollProfile(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.employeeService.DeletePayrollProfile(ctx.Request.Context(), id); err != nil {
		res := utils.BuildResponseFailed("failed delete payroll profile", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("success delete payroll profile", nil)
	ctx.JSON(http.StatusOK, res)
}
