package service

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/repository"
	"gorm.io/gorm"
)

type RbacService interface {
}

type rbacService struct {
	rbacRepository repository.RbacRepository
	db                            *gorm.DB
}

func NewRbacService(
	rbacRepo repository.RbacRepository,
	db *gorm.DB,
) RbacService {
	return &rbacService{
		rbacRepository: rbacRepo,
		db:                            db,
	}
}
