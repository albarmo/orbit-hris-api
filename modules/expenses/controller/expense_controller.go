package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/modules/expenses/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/expenses/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExpenseController interface {
    // categories
    CreateCategory(ctx *gin.Context)
    GetCategories(ctx *gin.Context)
    GetCategory(ctx *gin.Context)
    UpdateCategory(ctx *gin.Context)
    DeleteCategory(ctx *gin.Context)

    // expenses
    CreateExpense(ctx *gin.Context)
    GetExpenses(ctx *gin.Context)
    GetExpense(ctx *gin.Context)
    UpdateExpense(ctx *gin.Context)
    DeleteExpense(ctx *gin.Context)
}

type expenseController struct{
    svc service.ExpenseService
}

func NewExpenseController(s service.ExpenseService) ExpenseController { return &expenseController{svc: s} }

func (c *expenseController) CreateCategory(ctx *gin.Context) {
    var req dto.ExpenseCategoryCreateRequest
    if err := ctx.ShouldBindJSON(&req); err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    res, err := c.svc.CreateCategory(ctx.Request.Context(), req)
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    ctx.JSON(http.StatusCreated, res)
}

func (c *expenseController) GetCategories(ctx *gin.Context) {
    res, err := c.svc.FindAllCategories(ctx.Request.Context())
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetCategory(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"}); return }
    res, err := c.svc.FindCategoryByID(ctx.Request.Context(), id)
    if err != nil { ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
    ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) UpdateCategory(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"}); return }
    var req dto.ExpenseCategoryCreateRequest
    if err := ctx.ShouldBindJSON(&req); err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    res, err := c.svc.UpdateCategory(ctx.Request.Context(), id, req)
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) DeleteCategory(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"}); return }
    if err := c.svc.DeleteCategory(ctx.Request.Context(), id); err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    ctx.Status(http.StatusNoContent)
}

func (c *expenseController) CreateExpense(ctx *gin.Context) {
    var req dto.ExpenseCreateRequest
    if err := ctx.ShouldBindJSON(&req); err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    res, err := c.svc.CreateExpense(ctx.Request.Context(), req)
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    ctx.JSON(http.StatusCreated, res)
}

func (c *expenseController) GetExpenses(ctx *gin.Context) {
    res, err := c.svc.FindAllExpenses(ctx.Request.Context())
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpense(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"}); return }
    res, err := c.svc.FindExpenseByID(ctx.Request.Context(), id)
    if err != nil { ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
    ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) UpdateExpense(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"}); return }
    var req dto.ExpenseUpdateRequest
    if err := ctx.ShouldBindJSON(&req); err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
    res, err := c.svc.UpdateExpense(ctx.Request.Context(), id, req)
    if err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) DeleteExpense(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid id"}); return }
    if err := c.svc.DeleteExpense(ctx.Request.Context(), id); err != nil { ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
    ctx.Status(http.StatusNoContent)
}
