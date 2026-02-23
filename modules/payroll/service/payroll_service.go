package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PayrollService interface {
	Create(ctx context.Context, req dto.PayrollCreateRequest) (dto.PayrollResponse, error)
	FindAll(ctx context.Context, offset, limit int) ([]dto.PayrollResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (dto.PayrollResponse, error)
	Update(ctx context.Context, id uuid.UUID, req dto.PayrollUpdateRequest) (dto.PayrollResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type payrollService struct {
	payrollRepository repository.PayrollRepository
	db                            *gorm.DB
}

func NewPayrollService(
	payrollRepo repository.PayrollRepository,
	db *gorm.DB,
) PayrollService {
	return &payrollService{
		payrollRepository: payrollRepo,
		db:                            db,
	}
}

func (s *payrollService) Create(ctx context.Context, req dto.PayrollCreateRequest) (dto.PayrollResponse, error) {
	// map request to entity
	employeeID, _ := uuid.Parse(req.EmployeeID)
	periodID, _ := uuid.Parse(req.PayrollPeriodID)

	p := &entities.Payroll{
		EmployeeID:      employeeID,
		PayrollPeriodID: periodID,
		BasicSalary:     req.BasicSalary,
		TotalAllowance:  req.TotalAllowance,
		TotalDeduction:  req.TotalDeduction,
	}
	p.NetSalary = p.BasicSalary + p.TotalAllowance - p.TotalDeduction

	if err := s.payrollRepository.Create(ctx, p); err != nil {
		return dto.PayrollResponse{}, err
	}

	return dto.PayrollResponse{
		ID:              p.ID.String(),
		EmployeeID:      p.EmployeeID.String(),
		PayrollPeriodID: p.PayrollPeriodID.String(),
		BasicSalary:     p.BasicSalary,
		TotalAllowance:  p.TotalAllowance,
		TotalDeduction:  p.TotalDeduction,
		NetSalary:       p.NetSalary,
		GeneratedAt:     p.GeneratedAt.String(),
	}, nil
}

func (s *payrollService) FindAll(ctx context.Context, offset, limit int) ([]dto.PayrollResponse, error) {
	list, err := s.payrollRepository.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	var out []dto.PayrollResponse
	for _, p := range list {
		out = append(out, dto.PayrollResponse{
			ID:              p.ID.String(),
			EmployeeID:      p.EmployeeID.String(),
			PayrollPeriodID: p.PayrollPeriodID.String(),
			BasicSalary:     p.BasicSalary,
			TotalAllowance:  p.TotalAllowance,
			TotalDeduction:  p.TotalDeduction,
			NetSalary:       p.NetSalary,
			GeneratedAt:     p.GeneratedAt.String(),
		})
	}
	return out, nil
}

func (s *payrollService) FindByID(ctx context.Context, id uuid.UUID) (dto.PayrollResponse, error) {
	p, err := s.payrollRepository.FindByID(ctx, id)
	if err != nil {
		return dto.PayrollResponse{}, err
	}
	return dto.PayrollResponse{
		ID:              p.ID.String(),
		EmployeeID:      p.EmployeeID.String(),
		PayrollPeriodID: p.PayrollPeriodID.String(),
		BasicSalary:     p.BasicSalary,
		TotalAllowance:  p.TotalAllowance,
		TotalDeduction:  p.TotalDeduction,
		NetSalary:       p.NetSalary,
		GeneratedAt:     p.GeneratedAt.String(),
	}, nil
}

func (s *payrollService) Update(ctx context.Context, id uuid.UUID, req dto.PayrollUpdateRequest) (dto.PayrollResponse, error) {
	p, err := s.payrollRepository.FindByID(ctx, id)
	if err != nil {
		return dto.PayrollResponse{}, err
	}
	if req.BasicSalary != nil {
		p.BasicSalary = *req.BasicSalary
	}
	if req.TotalAllowance != nil {
		p.TotalAllowance = *req.TotalAllowance
	}
	if req.TotalDeduction != nil {
		p.TotalDeduction = *req.TotalDeduction
	}
	if req.PayrollPeriodID != nil {
		if v, err := uuid.Parse(*req.PayrollPeriodID); err == nil {
			p.PayrollPeriodID = v
		}
	}
	// recalc net salary
	p.NetSalary = p.BasicSalary + p.TotalAllowance - p.TotalDeduction

	if err := s.payrollRepository.Update(ctx, p); err != nil {
		return dto.PayrollResponse{}, err
	}

	return s.FindByID(ctx, id)
}

func (s *payrollService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.payrollRepository.Delete(ctx, id)
}
