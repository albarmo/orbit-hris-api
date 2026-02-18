package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(ctx context.Context, tx *gorm.DB, employee entities.Employee, personalInfo entities.EmployeePersonalInfo, address []entities.EmployeeAddress, legalInfo entities.EmployeeLegalInfo, payrollInfo entities.EmployeePayrollProfile) (entities.Employee, error)
	FindAll(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Employee], error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Employee, error)
	Update(ctx context.Context, tx *gorm.DB, employee entities.Employee) (entities.Employee, error)
	Delete(ctx context.Context, tx *gorm.DB, id uuid.UUID) error

	// Child table operations
	CreatePersonalInfo(ctx context.Context, tx *gorm.DB, info entities.EmployeePersonalInfo) (entities.EmployeePersonalInfo, error)
	GetPersonalInfoByEmployeeID(ctx context.Context, db *gorm.DB, employeeID uuid.UUID) (entities.EmployeePersonalInfo, error)
	UpdatePersonalInfo(ctx context.Context, tx *gorm.DB, info entities.EmployeePersonalInfo) (entities.EmployeePersonalInfo, error)
	DeletePersonalInfo(ctx context.Context, tx *gorm.DB, employeeID uuid.UUID) error

	CreateAddress(ctx context.Context, tx *gorm.DB, addr entities.EmployeeAddress) (entities.EmployeeAddress, error)
	GetAddressesByEmployeeID(ctx context.Context, db *gorm.DB, employeeID uuid.UUID) ([]entities.EmployeeAddress, error)
	GetAddressByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.EmployeeAddress, error)
	UpdateAddress(ctx context.Context, tx *gorm.DB, addr entities.EmployeeAddress) (entities.EmployeeAddress, error)
	DeleteAddress(ctx context.Context, tx *gorm.DB, id uuid.UUID) error

	CreateLegalInfo(ctx context.Context, tx *gorm.DB, info entities.EmployeeLegalInfo) (entities.EmployeeLegalInfo, error)
	GetLegalInfoByEmployeeID(ctx context.Context, db *gorm.DB, employeeID uuid.UUID) (entities.EmployeeLegalInfo, error)
	UpdateLegalInfo(ctx context.Context, tx *gorm.DB, info entities.EmployeeLegalInfo) (entities.EmployeeLegalInfo, error)
	DeleteLegalInfo(ctx context.Context, tx *gorm.DB, employeeID uuid.UUID) error

	CreatePayrollProfile(ctx context.Context, tx *gorm.DB, profile entities.EmployeePayrollProfile) (entities.EmployeePayrollProfile, error)
	GetPayrollByEmployeeID(ctx context.Context, db *gorm.DB, employeeID uuid.UUID) (entities.EmployeePayrollProfile, error)
	UpdatePayrollProfile(ctx context.Context, tx *gorm.DB, profile entities.EmployeePayrollProfile) (entities.EmployeePayrollProfile, error)
	DeletePayrollProfile(ctx context.Context, tx *gorm.DB, employeeID uuid.UUID) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) Create(ctx context.Context, tx *gorm.DB, employee entities.Employee, personalInfo entities.EmployeePersonalInfo, addresses []entities.EmployeeAddress, legalInfo entities.EmployeeLegalInfo, payrollInfo entities.EmployeePayrollProfile) (entities.Employee, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&employee).Error; err != nil {
		return entities.Employee{}, err
	}

	personalInfo.EmployeeID = employee.ID
	if err := tx.WithContext(ctx).Create(&personalInfo).Error; err != nil {
		return entities.Employee{}, err
	}

	for i := range addresses {
		addresses[i].EmployeeID = employee.ID
	}
	if err := tx.WithContext(ctx).Create(&addresses).Error; err != nil {
		return entities.Employee{}, err
	}

	legalInfo.EmployeeID = employee.ID
	if err := tx.WithContext(ctx).Create(&legalInfo).Error; err != nil {
		return entities.Employee{}, err
	}

	payrollInfo.EmployeeID = employee.ID
	if err := tx.WithContext(ctx).Create(&payrollInfo).Error; err != nil {
		return entities.Employee{}, err
	}

	return employee, nil
}

func (r *employeeRepository) FindAll(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Employee], error) {
	if db == nil {
		db = r.db
	}

	var employees []entities.Employee
	var page pagination.Page[entities.Employee]

	paginator, err := pagination.NewPaginator(db.WithContext(ctx).Model(&entities.Employee{}), filter)
	if err != nil {
		return nil, err
	}

	if err := paginator.Find(&employees).Error; err != nil {
		return nil, err
	}

	page.Set(employees, paginator.Page, paginator.Limit, paginator.Total)
	return &page, nil
}

func (r *employeeRepository) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Employee, error) {
	if db == nil {
		db = r.db
	}

	var employee entities.Employee
	if err := db.WithContext(ctx).Preload("User").Preload("Department").Preload("Position").Preload("Supervisor").Where("id = ?", id).First(&employee).Error; err != nil {
		return entities.Employee{}, err
	}

	return employee, nil
}

func (r *employeeRepository) Update(ctx context.Context, tx *gorm.DB, employee entities.Employee) (entities.Employee, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&employee).Error; err != nil {
		return entities.Employee{}, err
	}

	return employee, nil
}

func (r *employeeRepository) Delete(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entities.Employee{}).Error; err != nil {
		return err
	}
	return nil
}

// PersonalInfo
func (r *employeeRepository) CreatePersonalInfo(ctx context.Context, tx *gorm.DB, info entities.EmployeePersonalInfo) (entities.EmployeePersonalInfo, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&info).Error; err != nil {
		return entities.EmployeePersonalInfo{}, err
	}
	return info, nil
}

func (r *employeeRepository) GetPersonalInfoByEmployeeID(ctx context.Context, db *gorm.DB, employeeID uuid.UUID) (entities.EmployeePersonalInfo, error) {
	if db == nil {
		db = r.db
	}
	var info entities.EmployeePersonalInfo
	if err := db.WithContext(ctx).Where("employee_id = ?", employeeID).First(&info).Error; err != nil {
		return entities.EmployeePersonalInfo{}, err
	}
	return info, nil
}

func (r *employeeRepository) UpdatePersonalInfo(ctx context.Context, tx *gorm.DB, info entities.EmployeePersonalInfo) (entities.EmployeePersonalInfo, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.EmployeePersonalInfo{}).Where("employee_id = ?", info.EmployeeID).Updates(&info).Error; err != nil {
		return entities.EmployeePersonalInfo{}, err
	}
	return info, nil
}

func (r *employeeRepository) DeletePersonalInfo(ctx context.Context, tx *gorm.DB, employeeID uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("employee_id = ?", employeeID).Delete(&entities.EmployeePersonalInfo{}).Error; err != nil {
		return err
	}
	return nil
}

// Addresses
func (r *employeeRepository) CreateAddress(ctx context.Context, tx *gorm.DB, addr entities.EmployeeAddress) (entities.EmployeeAddress, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&addr).Error; err != nil {
		return entities.EmployeeAddress{}, err
	}
	return addr, nil
}

func (r *employeeRepository) GetAddressesByEmployeeID(ctx context.Context, db *gorm.DB, employeeID uuid.UUID) ([]entities.EmployeeAddress, error) {
	if db == nil {
		db = r.db
	}
	var addrs []entities.EmployeeAddress
	if err := db.WithContext(ctx).Where("employee_id = ?", employeeID).Find(&addrs).Error; err != nil {
		return nil, err
	}
	return addrs, nil
}

func (r *employeeRepository) GetAddressByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.EmployeeAddress, error) {
	if db == nil {
		db = r.db
	}
	var addr entities.EmployeeAddress
	if err := db.WithContext(ctx).Where("id = ?", id).First(&addr).Error; err != nil {
		return entities.EmployeeAddress{}, err
	}
	return addr, nil
}

func (r *employeeRepository) UpdateAddress(ctx context.Context, tx *gorm.DB, addr entities.EmployeeAddress) (entities.EmployeeAddress, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.EmployeeAddress{}).Where("id = ?", addr.ID).Updates(&addr).Error; err != nil {
		return entities.EmployeeAddress{}, err
	}
	return addr, nil
}

func (r *employeeRepository) DeleteAddress(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entities.EmployeeAddress{}).Error; err != nil {
		return err
	}
	return nil
}

// LegalInfo
func (r *employeeRepository) CreateLegalInfo(ctx context.Context, tx *gorm.DB, info entities.EmployeeLegalInfo) (entities.EmployeeLegalInfo, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&info).Error; err != nil {
		return entities.EmployeeLegalInfo{}, err
	}
	return info, nil
}

func (r *employeeRepository) GetLegalInfoByEmployeeID(ctx context.Context, db *gorm.DB, employeeID uuid.UUID) (entities.EmployeeLegalInfo, error) {
	if db == nil {
		db = r.db
	}
	var info entities.EmployeeLegalInfo
	if err := db.WithContext(ctx).Where("employee_id = ?", employeeID).First(&info).Error; err != nil {
		return entities.EmployeeLegalInfo{}, err
	}
	return info, nil
}

func (r *employeeRepository) UpdateLegalInfo(ctx context.Context, tx *gorm.DB, info entities.EmployeeLegalInfo) (entities.EmployeeLegalInfo, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.EmployeeLegalInfo{}).Where("employee_id = ?", info.EmployeeID).Updates(&info).Error; err != nil {
		return entities.EmployeeLegalInfo{}, err
	}
	return info, nil
}

func (r *employeeRepository) DeleteLegalInfo(ctx context.Context, tx *gorm.DB, employeeID uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("employee_id = ?", employeeID).Delete(&entities.EmployeeLegalInfo{}).Error; err != nil {
		return err
	}
	return nil
}

// Payroll
func (r *employeeRepository) CreatePayrollProfile(ctx context.Context, tx *gorm.DB, profile entities.EmployeePayrollProfile) (entities.EmployeePayrollProfile, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Create(&profile).Error; err != nil {
		return entities.EmployeePayrollProfile{}, err
	}
	return profile, nil
}

func (r *employeeRepository) GetPayrollByEmployeeID(ctx context.Context, db *gorm.DB, employeeID uuid.UUID) (entities.EmployeePayrollProfile, error) {
	if db == nil {
		db = r.db
	}
	var p entities.EmployeePayrollProfile
	if err := db.WithContext(ctx).Where("employee_id = ?", employeeID).First(&p).Error; err != nil {
		return entities.EmployeePayrollProfile{}, err
	}
	return p, nil
}

func (r *employeeRepository) UpdatePayrollProfile(ctx context.Context, tx *gorm.DB, profile entities.EmployeePayrollProfile) (entities.EmployeePayrollProfile, error) {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Model(&entities.EmployeePayrollProfile{}).Where("employee_id = ?", profile.EmployeeID).Updates(&profile).Error; err != nil {
		return entities.EmployeePayrollProfile{}, err
	}
	return profile, nil
}

func (r *employeeRepository) DeletePayrollProfile(ctx context.Context, tx *gorm.DB, employeeID uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}
	if err := tx.WithContext(ctx).Where("employee_id = ?", employeeID).Delete(&entities.EmployeePayrollProfile{}).Error; err != nil {
		return err
	}
	return nil
}
