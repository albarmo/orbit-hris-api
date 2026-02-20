package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/rbac/repository"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RbacService interface {
	// Roles
	CreateRole(ctx context.Context, tx *gorm.DB, r entities.Role) (entities.Role, error)
	FindRoles(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Role], error)
	GetRoleByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Role, error)
	UpdateRole(ctx context.Context, tx *gorm.DB, r entities.Role) (entities.Role, error)
	DeleteRole(ctx context.Context, tx *gorm.DB, id uuid.UUID) error

	// Permissions
	CreatePermission(ctx context.Context, tx *gorm.DB, p entities.Permission) (entities.Permission, error)
	FindPermissions(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Permission], error)
	GetPermissionByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Permission, error)
	UpdatePermission(ctx context.Context, tx *gorm.DB, p entities.Permission) (entities.Permission, error)
	DeletePermission(ctx context.Context, tx *gorm.DB, id uuid.UUID) error

	// User roles
	AssignRoleToUser(ctx context.Context, tx *gorm.DB, userRole entities.UserRole) error
	RemoveRoleFromUser(ctx context.Context, tx *gorm.DB, userID, roleID uuid.UUID) error
	GetRolesByUser(ctx context.Context, db *gorm.DB, userID uuid.UUID) ([]entities.Role, error)
}

type rbacService struct {
	rbacRepository repository.RbacRepository
	db             *gorm.DB
}

func NewRbacService(
	rbacRepo repository.RbacRepository,
	db *gorm.DB,
) RbacService {
	return &rbacService{rbacRepository: rbacRepo, db: db}
}

// Roles
func (s *rbacService) CreateRole(ctx context.Context, tx *gorm.DB, r entities.Role) (entities.Role, error) {
	return s.rbacRepository.CreateRole(ctx, tx, r)
}

func (s *rbacService) FindRoles(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Role], error) {
	return s.rbacRepository.FindRoles(ctx, db, filter)
}

func (s *rbacService) GetRoleByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Role, error) {
	return s.rbacRepository.GetRoleByID(ctx, db, id)
}

func (s *rbacService) UpdateRole(ctx context.Context, tx *gorm.DB, r entities.Role) (entities.Role, error) {
	return s.rbacRepository.UpdateRole(ctx, tx, r)
}

func (s *rbacService) DeleteRole(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	return s.rbacRepository.DeleteRole(ctx, tx, id)
}

// Permissions
func (s *rbacService) CreatePermission(ctx context.Context, tx *gorm.DB, p entities.Permission) (entities.Permission, error) {
	return s.rbacRepository.CreatePermission(ctx, tx, p)
}

func (s *rbacService) FindPermissions(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Permission], error) {
	return s.rbacRepository.FindPermissions(ctx, db, filter)
}

func (s *rbacService) GetPermissionByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Permission, error) {
	return s.rbacRepository.GetPermissionByID(ctx, db, id)
}

func (s *rbacService) UpdatePermission(ctx context.Context, tx *gorm.DB, p entities.Permission) (entities.Permission, error) {
	return s.rbacRepository.UpdatePermission(ctx, tx, p)
}

func (s *rbacService) DeletePermission(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	return s.rbacRepository.DeletePermission(ctx, tx, id)
}

// User roles
func (s *rbacService) AssignRoleToUser(ctx context.Context, tx *gorm.DB, userRole entities.UserRole) error {
	return s.rbacRepository.AssignRoleToUser(ctx, tx, userRole)
}

func (s *rbacService) RemoveRoleFromUser(ctx context.Context, tx *gorm.DB, userID, roleID uuid.UUID) error {
	return s.rbacRepository.RemoveRoleFromUser(ctx, tx, userID, roleID)
}

func (s *rbacService) GetRolesByUser(ctx context.Context, db *gorm.DB, userID uuid.UUID) ([]entities.Role, error) {
	return s.rbacRepository.GetRolesByUser(ctx, db, userID)
}
