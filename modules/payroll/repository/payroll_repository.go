package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PayrollRepository interface {
	Create(ctx context.Context, p *entities.Payroll) error
	FindAll(ctx context.Context, offset, limit int) ([]entities.Payroll, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Payroll, error)
	Update(ctx context.Context, p *entities.Payroll) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type payrollRepository struct {
	db *gorm.DB
}

func NewPayrollRepository(db *gorm.DB) PayrollRepository {
	return &payrollRepository{
		db: db,
	}
}

func (r *payrollRepository) Create(ctx context.Context, p *entities.Payroll) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *payrollRepository) FindAll(ctx context.Context, offset, limit int) ([]entities.Payroll, error) {
	var list []entities.Payroll
	q := r.db.WithContext(ctx).Preload("Employee").Preload("PayrollPeriod").Order("generated_at desc")
	if limit > 0 {
		q = q.Offset(offset).Limit(limit)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *payrollRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Payroll, error) {
	var p entities.Payroll
	if err := r.db.WithContext(ctx).Preload("Employee").Preload("PayrollPeriod").First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *payrollRepository) Update(ctx context.Context, p *entities.Payroll) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *payrollRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Payroll{}, "id = ?", id).Error
}
