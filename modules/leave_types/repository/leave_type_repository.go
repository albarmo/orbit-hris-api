package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeaveTypeRepository interface {
	Create(ctx context.Context, lt *entities.LeaveType) error
	FindAll(ctx context.Context) ([]entities.LeaveType, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.LeaveType, error)
	Update(ctx context.Context, lt *entities.LeaveType) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type leaveTypeRepository struct {
	db *gorm.DB
}

func NewLeaveTypeRepository(db *gorm.DB) LeaveTypeRepository {
	return &leaveTypeRepository{db: db}
}

func (r *leaveTypeRepository) Create(ctx context.Context, lt *entities.LeaveType) error {
	return r.db.WithContext(ctx).Create(lt).Error
}

func (r *leaveTypeRepository) FindAll(ctx context.Context) ([]entities.LeaveType, error) {
	var list []entities.LeaveType
	if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *leaveTypeRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.LeaveType, error) {
	var lt entities.LeaveType
	if err := r.db.WithContext(ctx).First(&lt, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &lt, nil
}

func (r *leaveTypeRepository) Update(ctx context.Context, lt *entities.LeaveType) error {
	return r.db.WithContext(ctx).Save(lt).Error
}

func (r *leaveTypeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.LeaveType{}, "id = ?", id).Error
}
