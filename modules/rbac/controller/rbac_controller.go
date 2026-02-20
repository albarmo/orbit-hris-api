package controller

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/service"
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/validation"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	RbacController interface {
	}

	rbacController struct {
		rbacService    service.RbacService
		rbacValidation *validation.RbacValidation
		db                             *gorm.DB
	}
)

func NewRbacController(injector *do.Injector, s service.RbacService) RbacController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	rbacValidation := validation.NewRbacValidation()
	return &rbacController{
		rbacService:    s,
		rbacValidation: rbacValidation,
		db:                             db,
	}
}
