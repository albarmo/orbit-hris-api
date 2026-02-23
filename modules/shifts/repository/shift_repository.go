package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShiftRepository interface {
    Create(ctx context.Context, s *entities.Shift) error
    FindAll(ctx context.Context) ([]entities.Shift, error)
    FindByID(ctx context.Context, id uuid.UUID) (*entities.Shift, error)
    Update(ctx context.Context, s *entities.Shift) error
    Delete(ctx context.Context, id uuid.UUID) error

    CreateEmployeeShift(ctx context.Context, es *entities.EmployeeShift) error
    FindEmployeeShiftsByEmployee(ctx context.Context, employeeID uuid.UUID) ([]entities.EmployeeShift, error)
    DeleteEmployeeShift(ctx context.Context, id uuid.UUID) error
}

type shiftRepository struct{
    db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) ShiftRepository {
    return &shiftRepository{db: db}
}

func (r *shiftRepository) Create(ctx context.Context, s *entities.Shift) error {
    return r.db.WithContext(ctx).Create(s).Error
}

func (r *shiftRepository) FindAll(ctx context.Context) ([]entities.Shift, error) {
    var list []entities.Shift
    if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
        return nil, err
    }
    return list, nil
}

func (r *shiftRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Shift, error) {
    var s entities.Shift
    if err := r.db.WithContext(ctx).First(&s, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &s, nil
}

func (r *shiftRepository) Update(ctx context.Context, s *entities.Shift) error {
    return r.db.WithContext(ctx).Save(s).Error
}

func (r *shiftRepository) Delete(ctx context.Context, id uuid.UUID) error {
    return r.db.WithContext(ctx).Delete(&entities.Shift{}, "id = ?", id).Error
}

func (r *shiftRepository) CreateEmployeeShift(ctx context.Context, es *entities.EmployeeShift) error {
    return r.db.WithContext(ctx).Create(es).Error
}

func (r *shiftRepository) FindEmployeeShiftsByEmployee(ctx context.Context, employeeID uuid.UUID) ([]entities.EmployeeShift, error) {
    var list []entities.EmployeeShift
    if err := r.db.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&list).Error; err != nil {
        return nil, err
    }
    return list, nil
}

func (r *shiftRepository) DeleteEmployeeShift(ctx context.Context, id uuid.UUID) error {
    return r.db.WithContext(ctx).Delete(&entities.EmployeeShift{}, "id = ?", id).Error
}
