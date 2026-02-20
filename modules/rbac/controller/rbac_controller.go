package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	RbacController interface {
	// Roles
	CreateRole(ctx *gin.Context)
	GetRoles(ctx *gin.Context)
	GetRoleByID(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)

	// Permissions
	CreatePermission(ctx *gin.Context)
	GetPermissions(ctx *gin.Context)
	GetPermissionByID(ctx *gin.Context)
	UpdatePermission(ctx *gin.Context)
	DeletePermission(ctx *gin.Context)

	// User-Role
	AssignRole(ctx *gin.Context)
	RemoveRole(ctx *gin.Context)
	GetRolesByUser(ctx *gin.Context)
	}

	rbacController struct {
		rbacService    service.RbacService
		rbacValidation *validation.RbacValidation
		db             *gorm.DB
	}
)

func NewRbacController(injector *do.Injector, s service.RbacService) RbacController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	rbacValidation := validation.NewRbacValidation()
	return &rbacController{
		rbacService:    s,
		rbacValidation: rbacValidation,
		db:             db,
	}
}

// Roles
func (c *rbacController) CreateRole(ctx *gin.Context) {
	var req dto.RoleCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err := c.rbacValidation.ValidateRoleCreate(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// build entity
	r := struct {
		Name        string
		Description string
	}{Name: req.Name, Description: req.Description}

	// delegate to service
	// convert to entities.Role inline to avoid import cycle
	result, err := c.rbacService.CreateRole(ctx.Request.Context(), nil, structToRole(r))
	if err != nil {
		res := utils.BuildResponseFailed("failed create role", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success create role", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *rbacController) GetRoles(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)
	page, err := c.rbacService.FindRoles(ctx.Request.Context(), nil, &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get roles", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

func (c *rbacController) GetRoleByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.rbacService.GetRoleByID(ctx.Request.Context(), nil, id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get role", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *rbacController) UpdateRole(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var req dto.RoleUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err := c.rbacValidation.ValidateRoleUpdate(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	rt := struct{ ID uuid.UUID; Name, Description string }{ID: id, Name: req.Name, Description: req.Description}
	result, err := c.rbacService.UpdateRole(ctx.Request.Context(), nil, structToRole(rt))
	if err != nil {
		res := utils.BuildResponseFailed("failed update role", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success update role", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *rbacController) DeleteRole(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	if err := c.rbacService.DeleteRole(ctx.Request.Context(), nil, id); err != nil {
		res := utils.BuildResponseFailed("failed delete role", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success delete role", nil)
	ctx.JSON(http.StatusOK, res)
}

// Permissions
func (c *rbacController) CreatePermission(ctx *gin.Context) {
	var req dto.PermissionCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err := c.rbacValidation.ValidatePermissionCreate(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	p := struct{ Name, Description string }{Name: req.Name, Description: req.Description}
	result, err := c.rbacService.CreatePermission(ctx.Request.Context(), nil, structToPermission(p))
	if err != nil {
		res := utils.BuildResponseFailed("failed create permission", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success create permission", result)
	ctx.JSON(http.StatusCreated, res)
}

func (c *rbacController) GetPermissions(ctx *gin.Context) {
	var filter = pagination.Filter{}
	filter.Bind(ctx)
	page, err := c.rbacService.FindPermissions(ctx.Request.Context(), nil, &filter)
	if err != nil {
		res := utils.BuildResponseFailed("failed get permissions", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success", page)
	ctx.JSON(http.StatusOK, res)
}

func (c *rbacController) GetPermissionByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.rbacService.GetPermissionByID(ctx.Request.Context(), nil, id)
	if err != nil {
		res := utils.BuildResponseFailed("failed get permission", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *rbacController) UpdatePermission(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var req dto.PermissionUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err := c.rbacValidation.ValidatePermissionUpdate(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	pt := struct{ ID uuid.UUID; Name, Description string }{ID: id, Name: req.Name, Description: req.Description}
	result, err := c.rbacService.UpdatePermission(ctx.Request.Context(), nil, structToPermission(pt))
	if err != nil {
		res := utils.BuildResponseFailed("failed update permission", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success update permission", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *rbacController) DeletePermission(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	if err := c.rbacService.DeletePermission(ctx.Request.Context(), nil, id); err != nil {
		res := utils.BuildResponseFailed("failed delete permission", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success delete permission", nil)
	ctx.JSON(http.StatusOK, res)
}

// User-Role
func (c *rbacController) AssignRole(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid user ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var req dto.AssignRoleRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if err := c.rbacValidation.ValidateAssignRole(req); err != nil {
		res := utils.BuildResponseFailed("validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	roleID, err := uuid.Parse(req.RoleID)
	if err != nil {
		res := utils.BuildResponseFailed("Invalid role ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	ur := struct{ UserID, RoleID uuid.UUID }{UserID: userID, RoleID: roleID}
	if err := c.rbacService.AssignRoleToUser(ctx.Request.Context(), nil, structToUserRole(ur)); err != nil {
		res := utils.BuildResponseFailed("failed assign role", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success assign role", nil)
	ctx.JSON(http.StatusCreated, res)
}

func (c *rbacController) RemoveRole(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid user ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	roleID, err := uuid.Parse(ctx.Param("role_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid role ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	if err := c.rbacService.RemoveRoleFromUser(ctx.Request.Context(), nil, userID, roleID); err != nil {
		res := utils.BuildResponseFailed("failed remove role", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success remove role", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *rbacController) GetRolesByUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_id"))
	if err != nil {
		res := utils.BuildResponseFailed("Invalid user ID", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	roles, err := c.rbacService.GetRolesByUser(ctx.Request.Context(), nil, userID)
	if err != nil {
		res := utils.BuildResponseFailed("failed get roles by user", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("success", roles)
	ctx.JSON(http.StatusOK, res)
}

// helpers to avoid import cycles: construct entities inline
func structToRole(src interface{}) entities.Role {
	// use type switches for the few inline structs we used
	switch v := src.(type) {
	case struct{ Name, Description string }:
		return entities.Role{Name: v.Name, Description: v.Description}
	case struct{ ID uuid.UUID; Name, Description string }:
		return entities.Role{ID: v.ID, Name: v.Name, Description: v.Description}
	}
	return entities.Role{}
}

func structToPermission(src interface{}) entities.Permission {
	switch v := src.(type) {
	case struct{ Name, Description string }:
		return entities.Permission{Name: v.Name, Description: v.Description}
	case struct{ ID uuid.UUID; Name, Description string }:
		return entities.Permission{ID: v.ID, Name: v.Name, Description: v.Description}
	}
	return entities.Permission{}
}

func structToUserRole(src interface{}) entities.UserRole {
	switch v := src.(type) {
	case struct{ UserID, RoleID uuid.UUID }:
		return entities.UserRole{UserID: v.UserID, RoleID: v.RoleID}
	}
	return entities.UserRole{}
}
