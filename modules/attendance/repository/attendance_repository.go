package repository

import (
	"context"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendanceRepository interface {
	FindAll(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Attendance], error)
	FindByID(id uuid.UUID) (*entities.Attendance, error)
	FindTodayByEmployeeID(employeeID uuid.UUID) (*entities.Attendance, error)
	Create(attendance *entities.Attendance) (*entities.Attendance, error)
	Update(attendance *entities.Attendance) (*entities.Attendance, error)
	Delete(id uuid.UUID) error
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{
		db: db,
	}
}

func (r *attendanceRepository) FindByID(id uuid.UUID) (*entities.Attendance, error) {
	var attendance entities.Attendance
	if err := r.db.Preload("Employee").Preload("Location").Where("id = ?", id).First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *attendanceRepository) FindTodayByEmployeeID(employeeID uuid.UUID) (*entities.Attendance, error) {
	var attendance entities.Attendance
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	err := r.db.Preload("Employee").Preload("Location").
		Where("employee_id = ?", employeeID).
		Where("check_in_time >= ? AND check_in_time < ?", today, tomorrow).
		First(&attendance).Error

	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *attendanceRepository) FindAll(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Attendance], error) {
	if db == nil {
		db = r.db
	}

	var attendances []entities.Attendance
	var page pagination.Page[entities.Attendance]

	paginator, err := pagination.NewPaginator(db.WithContext(ctx).Model(&entities.Attendance{}).Preload("Employee").Preload("Location"), filter)
	if err != nil {
		return nil, err
	}

	if err := paginator.Find(&attendances).Error; err != nil {
		return nil, err
	}

	page.Set(attendances, paginator.Page, paginator.Limit, paginator.Total)
	return &page, nil
}

func (r *attendanceRepository) Create(attendance *entities.Attendance) (*entities.Attendance, error) {
	if err := r.db.Create(attendance).Error; err != nil {
		return nil, err
	}
	// Preload the associated Employee and Location after creation
	if err := r.db.Preload("Employee").Preload("Location").First(&attendance, "id = ?", attendance.ID).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}

func (r *attendanceRepository) Update(attendance *entities.Attendance) (*entities.Attendance, error) {
	if err := r.db.Save(attendance).Error; err != nil {
		return nil, err
	}
	// Preload the associated Employee and Location after update
	if err := r.db.Preload("Employee").Preload("Location").First(&attendance, "id = ?", attendance.ID).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}

func (r *attendanceRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&entities.Attendance{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
