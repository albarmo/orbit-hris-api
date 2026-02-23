package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave_types/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave_types/repository"
	"github.com/google/uuid"
)

type LeaveTypeService interface {
    Create(ctx context.Context, req dto.LeaveTypeCreateRequest) (dto.LeaveTypeResponse, error)
    FindAll(ctx context.Context) ([]dto.LeaveTypeResponse, error)
    FindByID(ctx context.Context, id uuid.UUID) (dto.LeaveTypeResponse, error)
    Update(ctx context.Context, id uuid.UUID, req dto.LeaveTypeUpdateRequest) (dto.LeaveTypeResponse, error)
    Delete(ctx context.Context, id uuid.UUID) error
}

type leaveTypeService struct{
    repo repository.LeaveTypeRepository
}

func NewLeaveTypeService(r repository.LeaveTypeRepository) LeaveTypeService {
    return &leaveTypeService{repo: r}
}

func (s *leaveTypeService) Create(ctx context.Context, req dto.LeaveTypeCreateRequest) (dto.LeaveTypeResponse, error) {
    lt := &entities.LeaveType{
        Name: req.Name,
        MaxDaysPerYear: req.MaxDaysPerYear,
        IsPaid: req.IsPaid,
    }
    if err := s.repo.Create(ctx, lt); err != nil {
        return dto.LeaveTypeResponse{}, err
    }
    return dto.LeaveTypeResponse{ID: lt.ID.String(), Name: lt.Name, MaxDaysPerYear: lt.MaxDaysPerYear, IsPaid: lt.IsPaid}, nil
}

func (s *leaveTypeService) FindAll(ctx context.Context) ([]dto.LeaveTypeResponse, error) {
    list, err := s.repo.FindAll(ctx)
    if err != nil {
        return nil, err
    }
    var out []dto.LeaveTypeResponse
    for _, lt := range list {
        out = append(out, dto.LeaveTypeResponse{ID: lt.ID.String(), Name: lt.Name, MaxDaysPerYear: lt.MaxDaysPerYear, IsPaid: lt.IsPaid})
    }
    return out, nil
}

func (s *leaveTypeService) FindByID(ctx context.Context, id uuid.UUID) (dto.LeaveTypeResponse, error) {
    lt, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return dto.LeaveTypeResponse{}, err
    }
    return dto.LeaveTypeResponse{ID: lt.ID.String(), Name: lt.Name, MaxDaysPerYear: lt.MaxDaysPerYear, IsPaid: lt.IsPaid}, nil
}

func (s *leaveTypeService) Update(ctx context.Context, id uuid.UUID, req dto.LeaveTypeUpdateRequest) (dto.LeaveTypeResponse, error) {
    lt, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return dto.LeaveTypeResponse{}, err
    }
    if req.Name != nil { lt.Name = *req.Name }
    if req.MaxDaysPerYear != nil { lt.MaxDaysPerYear = *req.MaxDaysPerYear }
    if req.IsPaid != nil { lt.IsPaid = *req.IsPaid }
    if err := s.repo.Update(ctx, lt); err != nil {
        return dto.LeaveTypeResponse{}, err
    }
    return s.FindByID(ctx, id)
}

func (s *leaveTypeService) Delete(ctx context.Context, id uuid.UUID) error {
    return s.repo.Delete(ctx, id)
}
