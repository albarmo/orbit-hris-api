package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/leave_types/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave_types/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LeaveTypeController interface {
    Create(ctx *gin.Context)
    GetAll(ctx *gin.Context)
    GetByID(ctx *gin.Context)
    Update(ctx *gin.Context)
    Delete(ctx *gin.Context)
}

type leaveTypeController struct{
    svc service.LeaveTypeService
}

func NewLeaveTypeController(s service.LeaveTypeService) LeaveTypeController {
    return &leaveTypeController{svc: s}
}

func (c *leaveTypeController) Create(ctx *gin.Context) {
    var req dto.LeaveTypeCreateRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    res, err := c.svc.Create(ctx.Request.Context(), req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, res)
}

func (c *leaveTypeController) GetAll(ctx *gin.Context) {
    res, err := c.svc.FindAll(ctx.Request.Context())
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, res)
}

func (c *leaveTypeController) GetByID(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    res, err := c.svc.FindByID(ctx.Request.Context(), id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, res)
}

func (c *leaveTypeController) Update(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    var req dto.LeaveTypeUpdateRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    res, err := c.svc.Update(ctx.Request.Context(), id, req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, res)
}

func (c *leaveTypeController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    if err := c.svc.Delete(ctx.Request.Context(), id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.Status(http.StatusNoContent)
}
