package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RbacRepository interface {
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

	// User-Role assignment
	AssignRoleToUser(ctx context.Context, tx *gorm.DB, userRole entities.UserRole) error
	RemoveRoleFromUser(ctx context.Context, tx *gorm.DB, userID, roleID uuid.UUID) error
	GetRolesByUser(ctx context.Context, db *gorm.DB, userID uuid.UUID) ([]entities.Role, error)
}

type rbacRepository struct {
	db *gorm.DB
}

func NewRbacRepository(db *gorm.DB) RbacRepository {
	return &rbacRepository{db: db}
}

// Roles
func (r *rbacRepository) CreateRole(ctx context.Context, tx *gorm.DB, role entities.Role) (entities.Role, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&role).Error; err != nil {
		return entities.Role{}, err
	}
	return role, nil
}

func (r *rbacRepository) FindRoles(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Role], error) {
	if db == nil {
		db = r.db
	}
	var items []entities.Role
	var page pagination.Page[entities.Role]
	paginator, err := pagination.NewPaginator(db.WithContext(ctx).Model(&entities.Role{}), filter)
	if err != nil {
		return nil, err
	}
	if err := paginator.Find(&items).Error; err != nil {
		return nil, err
	}
	page.Set(items, paginator.Page, paginator.Limit, paginator.Total)
	return &page, nil
}

func (r *rbacRepository) GetRoleByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Role, error) {
	if db == nil {
		db = r.db
	}
	var role entities.Role
	if err := db.WithContext(ctx).Where("id = ?", id).First(&role).Error; err != nil {
		return entities.Role{}, err
	}
	return role, nil
}

func (r *rbacRepository) UpdateRole(ctx context.Context, tx *gorm.DB, role entities.Role) (entities.Role, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.Role{}).Where("id = ?", role.ID).Updates(&role).Error; err != nil {
		return entities.Role{}, err
	}
	return role, nil
}

func (r *rbacRepository) DeleteRole(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entities.Role{}).Error; err != nil {
		return err
	}
	return nil
}

// Permissions
func (r *rbacRepository) CreatePermission(ctx context.Context, tx *gorm.DB, p entities.Permission) (entities.Permission, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&p).Error; err != nil {
		return entities.Permission{}, err
	}
	return p, nil
}

func (r *rbacRepository) FindPermissions(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Permission], error) {
	if db == nil {
		db = r.db
	}
	var items []entities.Permission
	var page pagination.Page[entities.Permission]
	paginator, err := pagination.NewPaginator(db.WithContext(ctx).Model(&entities.Permission{}), filter)
	if err != nil {
		return nil, err
	}
	if err := paginator.Find(&items).Error; err != nil {
		return nil, err
	}
	page.Set(items, paginator.Page, paginator.Limit, paginator.Total)
	return &page, nil
}

func (r *rbacRepository) GetPermissionByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Permission, error) {
	if db == nil {
		db = r.db
	}
	var p entities.Permission
	if err := db.WithContext(ctx).Where("id = ?", id).First(&p).Error; err != nil {
		return entities.Permission{}, err
	}
	return p, nil
}

func (r *rbacRepository) UpdatePermission(ctx context.Context, tx *gorm.DB, p entities.Permission) (entities.Permission, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.Permission{}).Where("id = ?", p.ID).Updates(&p).Error; err != nil {
		return entities.Permission{}, err
	}
	return p, nil
}

func (r *rbacRepository) DeletePermission(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entities.Permission{}).Error; err != nil {
		return err
	}
	return nil
}

// User-Role assignment
func (r *rbacRepository) AssignRoleToUser(ctx context.Context, tx *gorm.DB, userRole entities.UserRole) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&userRole).Error; err != nil {
		return err
	}
	return nil
}

func (r *rbacRepository) RemoveRoleFromUser(ctx context.Context, tx *gorm.DB, userID, roleID uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&entities.UserRole{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *rbacRepository) GetRolesByUser(ctx context.Context, db *gorm.DB, userID uuid.UUID) ([]entities.Role, error) {
	if db == nil {
		db = r.db
	}
	var roles []entities.Role
	// join user_roles
	if err := db.WithContext(ctx).Model(&entities.Role{}).
		Preload("Permissions").
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id = ?", userID).
		Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
