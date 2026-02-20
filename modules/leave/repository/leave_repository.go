package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeaveRepository interface {
	FindAll(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Leave], error)
	FindByID(id uuid.UUID) (*entities.Leave, error)
	Create(leave *entities.Leave) (*entities.Leave, error)
	Update(leave *entities.Leave) (*entities.Leave, error)
	Delete(id uuid.UUID) error
}

type leaveRepository struct {
	db *gorm.DB
}

func NewLeaveRepository(db *gorm.DB) LeaveRepository {
	return &leaveRepository{
		db: db,
	}
}

func (r *leaveRepository) FindAll(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Leave], error) {
	if db == nil {
		db = r.db
	}

	var items []entities.Leave
	var page pagination.Page[entities.Leave]

	paginator, err := pagination.NewPaginator(db.WithContext(ctx).Model(&entities.Leave{}).Preload("Employee"), filter)
	if err != nil {
		return nil, err
	}

	if err := paginator.Find(&items).Error; err != nil {
		return nil, err
	}

	page.Set(items, paginator.Page, paginator.Limit, paginator.Total)
	return &page, nil
}

func (r *leaveRepository) FindByID(id uuid.UUID) (*entities.Leave, error) {
	var leave entities.Leave
	if err := r.db.Preload("Employee").Where("id = ?", id).First(&leave).Error; err != nil {
		return nil, err
	}
	return &leave, nil
}

func (r *leaveRepository) Create(leave *entities.Leave) (*entities.Leave, error) {
	if err := r.db.Create(leave).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Employee").First(leave, "id = ?", leave.ID).Error; err != nil {
		return nil, err
	}
	return leave, nil
}

func (r *leaveRepository) Update(leave *entities.Leave) (*entities.Leave, error) {
	if err := r.db.Save(leave).Error; err != nil {
		return nil, err
	}
	if err := r.db.Preload("Employee").First(leave, "id = ?", leave.ID).Error; err != nil {
		return nil, err
	}
	return leave, nil
}

func (r *leaveRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&entities.Leave{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
