package service

import (
	"context"
	"errors"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/leave/repository"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/pagination"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeaveService interface {
	FindAll(ctx context.Context, filter *pagination.Filter) (*pagination.Page[entities.Leave], error)
	GetByID(id string) (*entities.Leave, error)
	Create(req dto.LeaveCreateRequest) (*entities.Leave, error)
	Update(id string, req dto.LeaveUpdateRequest) (*entities.Leave, error)
	Delete(id string) error
}

type leaveService struct {
	leaveRepository repository.LeaveRepository
	db              *gorm.DB
}

func NewLeaveService(
	leaveRepo repository.LeaveRepository,
	db *gorm.DB,
) LeaveService {
	return &leaveService{
		leaveRepository: leaveRepo,
		db:              db,
	}
}

func (s *leaveService) FindAll(ctx context.Context, filter *pagination.Filter) (*pagination.Page[entities.Leave], error) {
	return s.leaveRepository.FindAll(ctx, s.db, filter)
}

func (s *leaveService) GetByID(id string) (*entities.Leave, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	return s.leaveRepository.FindByID(uid)
}

func (s *leaveService) Create(req dto.LeaveCreateRequest) (*entities.Leave, error) {
	leave := &entities.Leave{
		EmployeeID: req.EmployeeID,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Reason:     req.Reason,
		Status:     "pending",
	}
	return s.leaveRepository.Create(leave)
}

func (s *leaveService) Update(id string, req dto.LeaveUpdateRequest) (*entities.Leave, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	leave, err := s.leaveRepository.FindByID(uid)
	if err != nil {
		return nil, err
	}

	if req.Reason != "" {
		leave.Reason = req.Reason
	}
	if req.Status != "" {
		leave.Status = req.Status
	}

	return s.leaveRepository.Update(leave)
}

func (s *leaveService) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id")
	}
	return s.leaveRepository.Delete(uid)
}
