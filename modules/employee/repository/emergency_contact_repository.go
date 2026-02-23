package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmergencyContactRepository interface {
    Create(ctx context.Context, e *entities.EmployeeEmergencyContact) error
    FindByEmployee(ctx context.Context, employeeID uuid.UUID) ([]entities.EmployeeEmergencyContact, error)
    Delete(ctx context.Context, id uuid.UUID) error
}

type emergencyContactRepository struct{
    db *gorm.DB
}

func NewEmergencyContactRepository(db *gorm.DB) EmergencyContactRepository {
    return &emergencyContactRepository{db: db}
}

func (r *emergencyContactRepository) Create(ctx context.Context, e *entities.EmployeeEmergencyContact) error {
    return r.db.WithContext(ctx).Create(e).Error
}

func (r *emergencyContactRepository) FindByEmployee(ctx context.Context, employeeID uuid.UUID) ([]entities.EmployeeEmergencyContact, error) {
    var list []entities.EmployeeEmergencyContact
    if err := r.db.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&list).Error; err != nil {
        return nil, err
    }
    return list, nil
}

func (r *emergencyContactRepository) Delete(ctx context.Context, id uuid.UUID) error {
    return r.db.WithContext(ctx).Delete(&entities.EmployeeEmergencyContact{}, "id = ?", id).Error
}
