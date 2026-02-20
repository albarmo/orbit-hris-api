package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/master/repository"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterService interface {
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
	CreatePosition(ctx context.Context, tx *gorm.DB, pos entities.Position) (entities.Position, error)
	FindPositions(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Position], error)
	GetPositionByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Position, error)
	UpdatePosition(ctx context.Context, tx *gorm.DB, pos entities.Position) (entities.Position, error)
	DeletePosition(ctx context.Context, tx *gorm.DB, id uuid.UUID) error
}

type masterService struct {
	masterRepository repository.MasterRepository
	db               *gorm.DB
}

func NewMasterService(
	masterRepo repository.MasterRepository,
	db *gorm.DB,
) MasterService {
	return &masterService{
		masterRepository: masterRepo,
		db:               db,
	}
}

// Departments
func (s *masterService) CreateDepartment(ctx context.Context, tx *gorm.DB, dept entities.Department) (entities.Department, error) {
	return s.masterRepository.CreateDepartment(ctx, tx, dept)
}

func (s *masterService) FindDepartments(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Department], error) {
	return s.masterRepository.FindDepartments(ctx, db, filter)
}

func (s *masterService) GetDepartmentByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Department, error) {
	return s.masterRepository.GetDepartmentByID(ctx, db, id)
}

func (s *masterService) UpdateDepartment(ctx context.Context, tx *gorm.DB, dept entities.Department) (entities.Department, error) {
	return s.masterRepository.UpdateDepartment(ctx, tx, dept)
}

func (s *masterService) DeleteDepartment(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	return s.masterRepository.DeleteDepartment(ctx, tx, id)
}

// Locations
func (s *masterService) CreateLocation(ctx context.Context, tx *gorm.DB, loc entities.Location) (entities.Location, error) {
	return s.masterRepository.CreateLocation(ctx, tx, loc)
}

func (s *masterService) FindLocations(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Location], error) {
	return s.masterRepository.FindLocations(ctx, db, filter)
}

func (s *masterService) GetLocationByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Location, error) {
	return s.masterRepository.GetLocationByID(ctx, db, id)
}

func (s *masterService) UpdateLocation(ctx context.Context, tx *gorm.DB, loc entities.Location) (entities.Location, error) {
	return s.masterRepository.UpdateLocation(ctx, tx, loc)
}

func (s *masterService) DeleteLocation(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	return s.masterRepository.DeleteLocation(ctx, tx, id)
}

// Positions
func (s *masterService) CreatePosition(ctx context.Context, tx *gorm.DB, pos entities.Position) (entities.Position, error) {
	return s.masterRepository.CreatePosition(ctx, tx, pos)
}

func (s *masterService) FindPositions(ctx context.Context, db *gorm.DB, filter *pagination.Filter) (*pagination.Page[entities.Position], error) {
	return s.masterRepository.FindPositions(ctx, db, filter)
}

func (s *masterService) GetPositionByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (entities.Position, error) {
	return s.masterRepository.GetPositionByID(ctx, db, id)
}

func (s *masterService) UpdatePosition(ctx context.Context, tx *gorm.DB, pos entities.Position) (entities.Position, error) {
	return s.masterRepository.UpdatePosition(ctx, tx, pos)
}

func (s *masterService) DeletePosition(ctx context.Context, tx *gorm.DB, id uuid.UUID) error {
	return s.masterRepository.DeletePosition(ctx, tx, id)
}
