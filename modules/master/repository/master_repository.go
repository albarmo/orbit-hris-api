package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterRepository interface {
	// Departments
	CreateDepartment(ctx context.Context, tx *gorm.DB, dept entities.Department) (entities.Department, error)
	FindDepartments(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Department], error)
	GetDepartmentByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Department, error)
	UpdateDepartment(ctx context.Context, tx *gorm.DB, dept entities.Department) (entities.Department, error)
	DeleteDepartment(ctx context.Context, tx *gorm.DB, id uuid.UUID) error

	// Locations
	CreateLocation(ctx context.Context, tx *gorm.DB, loc entities.Location) (entities.Location, error)
	FindLocations(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Location], error)
	GetLocationByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Location, error)
	UpdateLocation(ctx context.Context, tx *gorm.DB, loc entities.Location) (entities.Location, error)
	DeleteLocation(ctx context.Context, tx *gorm.DB, id uuid.UUID) error

	// Positions
	CreatePosition(ctx context.Context, tx *gorm.DB, p entities.Position) (entities.Position, error)
	FindPositions(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Position], error)
	GetPositionByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Position, error)
	UpdatePosition(ctx context.Context, tx *gorm.DB, p entities.Position) (entities.Position, error)
	DeletePosition(ctx context.Context, tx *gorm.DB, id uuid.UUID) error
}

type masterRepository struct {
	db *gorm.DB
}

func NewMasterRepository(db *gorm.DB) MasterRepository {
	return &masterRepository{
		db: db,
	}
}

// Departments
func (r *masterRepository) CreateDepartment(ctx context.Context, tx *gorm.DB, dept entities.Department) (entities.Department, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&dept).Error; err != nil {
		return entities.Department{}, err
	}
	return dept, nil
}

func (r *masterRepository) FindDepartments(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Department], error) {
	if db == nil {
		db = r.db
	}
	var items []entities.Department
	var page pagination.Page[entities.Department]
	paginator, err := pagination.NewPaginator(db.WithContext(ctx).Model(&entities.Department{}), filter)
	if err != nil {
		return nil, err
	}
	if err := paginator.Find(&items).Error; err != nil {
		return nil, err
	}
	page.Set(items, paginator.Page, paginator.Limit, paginator.Total)
	return &page, nil
}

func (r *masterRepository) GetDepartmentByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Department, error) {
	if db == nil {
		db = r.db
	}
	var d entities.Department
	if err := db.WithContext(ctx).Where("id = ?", id).First(&d).Error; err != nil {
		return entities.Department{}, err
	}
	return d, nil
}

func (r *masterRepository) UpdateDepartment(ctx context.Context, tx *gorm.DB, dept entities.Department) (entities.Department, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.Department{}).Where("id = ?", dept.ID).Updates(&dept).Error; err != nil {
		return entities.Department{}, err
	}
	return dept, nil
}

func (r *masterRepository) DeleteDepartment(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entities.Department{}).Error; err != nil {
		return err
	}
	return nil
}

// Locations
func (r *masterRepository) CreateLocation(ctx context.Context, tx *gorm.DB, loc entities.Location) (entities.Location, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&loc).Error; err != nil {
		return entities.Location{}, err
	}
	return loc, nil
}

func (r *masterRepository) FindLocations(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Location], error) {
	if db == nil {
		db = r.db
	}
	var items []entities.Location
	var page pagination.Page[entities.Location]
	paginator, err := pagination.NewPaginator(db.WithContext(ctx).Model(&entities.Location{}), filter)
	if err != nil {
		return nil, err
	}
	if err := paginator.Find(&items).Error; err != nil {
		return nil, err
	}
	page.Set(items, paginator.Page, paginator.Limit, paginator.Total)
	return &page, nil
}

func (r *masterRepository) GetLocationByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Location, error) {
	if db == nil {
		db = r.db
	}
	var l entities.Location
	if err := db.WithContext(ctx).Where("id = ?", id).First(&l).Error; err != nil {
		return entities.Location{}, err
	}
	return l, nil
}

func (r *masterRepository) UpdateLocation(ctx context.Context, tx *gorm.DB, loc entities.Location) (entities.Location, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.Location{}).Where("id = ?", loc.ID).Updates(&loc).Error; err != nil {
		return entities.Location{}, err
	}
	return loc, nil
}

func (r *masterRepository) DeleteLocation(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entities.Location{}).Error; err != nil {
		return err
	}
	return nil
}

// Positions
func (r *masterRepository) CreatePosition(ctx context.Context, tx *gorm.DB, p entities.Position) (entities.Position, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&p).Error; err != nil {
		return entities.Position{}, err
	}
	return p, nil
}

func (r *masterRepository) FindPositions(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Position], error) {
	if db == nil {
		db = r.db
	}
	var items []entities.Position
	var page pagination.Page[entities.Position]
	paginator, err := pagination.NewPaginator(db.WithContext(ctx).Model(&entities.Position{}).Preload("Department"), filter)
	if err != nil {
		return nil, err
	}
	if err := paginator.Find(&items).Error; err != nil {
		return nil, err
	}
	page.Set(items, paginator.Page, paginator.Limit, paginator.Total)
	return &page, nil
}

func (r *masterRepository) GetPositionByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Position, error) {
	if db == nil {
		db = r.db
	}
	var p entities.Position
	if err := db.WithContext(ctx).Preload("Department").Where("id = ?", id).First(&p).Error; err != nil {
		return entities.Position{}, err
	}
	return p, nil
}

func (r *masterRepository) UpdatePosition(ctx context.Context, tx *gorm.DB, p entities.Position) (entities.Position, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.Position{}).Where("id = ?", p.ID).Updates(&p).Error; err != nil {
		return entities.Position{}, err
	}
	return p, nil
}

func (r *masterRepository) DeletePosition(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entities.Position{}).Error; err != nil {
		return err
	}
	return nil
}
