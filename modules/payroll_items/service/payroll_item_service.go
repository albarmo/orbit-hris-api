package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll_items/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/payroll_items/repository"
	"github.com/google/uuid"
)

type PayrollItemService interface {
	Create(ctx context.Context, req dto.PayrollItemCreateRequest) (dto.PayrollItemResponse, error)
	FindAll(ctx context.Context) ([]dto.PayrollItemResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (dto.PayrollItemResponse, error)
	Update(ctx context.Context, id uuid.UUID, req dto.PayrollItemUpdateRequest) (dto.PayrollItemResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type payrollItemService struct {
	repo repository.PayrollItemRepository
}

func NewPayrollItemService(r repository.PayrollItemRepository) PayrollItemService {
	return &payrollItemService{repo: r}
}

func (s *payrollItemService) Create(ctx context.Context, req dto.PayrollItemCreateRequest) (dto.PayrollItemResponse, error) {
	p := &entities.PayrollItem{PayrollID: uuid.MustParse(req.PayrollID), Name: req.Name, Amount: req.Amount}
	if err := s.repo.Create(ctx, p); err != nil {
		return dto.PayrollItemResponse{}, err
	}
	return dto.PayrollItemResponse{ID: p.ID.String(), PayrollID: p.PayrollID.String(), Name: p.Name, Amount: p.Amount}, nil
}

func (s *payrollItemService) FindAll(ctx context.Context) ([]dto.PayrollItemResponse, error) {
	list, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var out []dto.PayrollItemResponse
	for _, p := range list {
		out = append(out, dto.PayrollItemResponse{ID: p.ID.String(), PayrollID: p.PayrollID.String(), Name: p.Name, Amount: p.Amount})
	}
	return out, nil
}

func (s *payrollItemService) FindByID(ctx context.Context, id uuid.UUID) (dto.PayrollItemResponse, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.PayrollItemResponse{}, err
	}
	return dto.PayrollItemResponse{ID: p.ID.String(), PayrollID: p.PayrollID.String(), Name: p.Name, Amount: p.Amount}, nil
}

func (s *payrollItemService) Update(ctx context.Context, id uuid.UUID, req dto.PayrollItemUpdateRequest) (dto.PayrollItemResponse, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.PayrollItemResponse{}, err
	}
	if req.Name != nil {
		p.Name = *req.Name
	}
	if req.Amount != nil {
		p.Amount = *req.Amount
	}
	if err := s.repo.Update(ctx, p); err != nil {
		return dto.PayrollItemResponse{}, err
	}
	return s.FindByID(ctx, id)
}

func (s *payrollItemService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
