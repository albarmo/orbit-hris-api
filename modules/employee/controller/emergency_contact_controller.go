package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmergencyContactController interface {
    Create(ctx *gin.Context)
    GetByEmployee(ctx *gin.Context)
    Delete(ctx *gin.Context)
}

type emergencyContactController struct{
    svc service.EmergencyContactService
}

func NewEmergencyContactController(s service.EmergencyContactService) EmergencyContactController {
    return &emergencyContactController{svc: s}
}

func (c *emergencyContactController) Create(ctx *gin.Context) {
    employeeID, err := uuid.Parse(ctx.Param("id"))
    if err != nil {
        res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }
    var req dto.EmergencyContactCreateRequest
    if err := ctx.ShouldBind(&req); err != nil {
        res := utils.BuildResponseFailed("failed get data from body", err.Error(), nil)
        ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
        return
    }
    out, err := c.svc.Create(ctx.Request.Context(), employeeID, req)
    if err != nil {
        res := utils.BuildResponseFailed("failed create", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }
    res := utils.BuildResponseSuccess("success create", out)
    ctx.JSON(http.StatusCreated, res)
}

func (c *emergencyContactController) GetByEmployee(ctx *gin.Context) {
    employeeID, err := uuid.Parse(ctx.Param("id"))
    if err != nil {
        res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }
    out, err := c.svc.FindByEmployee(ctx.Request.Context(), employeeID)
    if err != nil {
        res := utils.BuildResponseFailed("failed get", err.Error(), nil)
        ctx.JSON(http.StatusBadRequest, res)
        return
    }
    res := utils.BuildResponseSuccess("success get", out)
    ctx.JSON(http.StatusOK, res)
}

func (c *emergencyContactController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("contact_id")
    id, err := uuid.Parse(idStr)
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
