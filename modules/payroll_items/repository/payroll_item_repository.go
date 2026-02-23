package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PayrollItemRepository interface {
	Create(ctx context.Context, p *entities.PayrollItem) error
	FindAll(ctx context.Context) ([]entities.PayrollItem, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.PayrollItem, error)
	Update(ctx context.Context, p *entities.PayrollItem) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type payrollItemRepository struct{ db *gorm.DB }

func NewPayrollItemRepository(db *gorm.DB) PayrollItemRepository {
	return &payrollItemRepository{db: db}
}

func (r *payrollItemRepository) Create(ctx context.Context, p *entities.PayrollItem) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *payrollItemRepository) FindAll(ctx context.Context) ([]entities.PayrollItem, error) {
	var list []entities.PayrollItem
	if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *payrollItemRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.PayrollItem, error) {
	var p entities.PayrollItem
	if err := r.db.WithContext(ctx).First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *payrollItemRepository) Update(ctx context.Context, p *entities.PayrollItem) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *payrollItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.PayrollItem{}, "id = ?", id).Error
}
