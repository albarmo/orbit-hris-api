package controller

import (
	"net/http"

	rbacService "github.com/Caknoooo/go-gin-clean-starter/modules/rbac/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/query"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/utils"
	"github.com/Caknoooo/go-pagination"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	UserController interface {
		Me(ctx *gin.Context)
		GetUserById(ctx *gin.Context)
		GetAllUser(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	userController struct {
		userService    service.UserService
		rbacService    rbacService.RbacService
		userValidation *validation.UserValidation
		db             *gorm.DB
	}
)

func NewUserController(injector *do.Injector, us service.UserService) UserController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	rs := do.MustInvoke[rbacService.RbacService](injector)
	userValidation := validation.NewUserValidation()
	return &userController{
		userService:    us,
		rbacService:    rs,
		userValidation: userValidation,
		db:             db,
	}
}

func (c *userController) GetAllUser(ctx *gin.Context) {
	var filter = &query.UserFilter{}
	filter.BindPagination(ctx)

	ctx.ShouldBindQuery(filter)

	users, total, err := pagination.PaginatedQueryWithIncludable[query.User](c.db, filter)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	paginationResponse := pagination.CalculatePagination(filter.Pagination, total)
	response := pagination.NewPaginatedResponse(http.StatusOK, dto.MESSAGE_SUCCESS_GET_LIST_USER, users, paginationResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")

	result, err := c.userService.GetUserById(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if uid, err := uuid.Parse(result.ID); err == nil {
		if roles, err := c.rbacService.GetRolesByUser(ctx.Request.Context(), nil, uid); err == nil {
			result.Roles = make([]dto.RoleResponse, 0)
			for _, r := range roles {
				perms := make([]dto.PermissionResponse, 0)
				for _, p := range r.Permissions {
					perms = append(perms, dto.PermissionResponse{
						ID:          p.ID.String(),
						Name:        p.Name,
						Description: p.Description,
					})
				}
				result.Roles = append(result.Roles, dto.RoleResponse{
					ID:          r.ID.String(),
					Name:        r.Name,
					Description: r.Description,
					Permissions: perms,
				})
			}
		}
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Me(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	result, err := c.userService.GetUserById(ctx.Request.Context(), userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if uid, err := uuid.Parse(result.ID); err == nil {
		if roles, err := c.rbacService.GetRolesByUser(ctx.Request.Context(), nil, uid); err == nil {
			result.Roles = make([]dto.RoleResponse, 0)
			for _, r := range roles {
				perms := make([]dto.PermissionResponse, 0)
				for _, p := range r.Permissions {
					perms = append(perms, dto.PermissionResponse{
						ID:          p.ID.String(),
						Name:        p.Name,
						Description: p.Description,
					})
				}
				result.Roles = append(result.Roles, dto.RoleResponse{
					ID:          r.ID.String(),
					Name:        r.Name,
					Description: r.Description,
					Permissions: perms,
				})
			}
		}
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Update(ctx *gin.Context) {
	var req dto.UserUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.userValidation.ValidateUserUpdateRequest(req); err != nil {
		res := utils.BuildResponseFailed("Validation failed", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userId := ctx.MustGet("user_id").(string)
	result, err := c.userService.Update(ctx.Request.Context(), req, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Delete(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	if err := c.userService.Delete(ctx.Request.Context(), userId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_USER, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_USER, nil)
	ctx.JSON(http.StatusOK, res)
}
