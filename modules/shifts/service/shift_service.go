package service

import (
	"context"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/shifts/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/shifts/repository"
	"github.com/google/uuid"
)

type ShiftService interface {
    Create(ctx context.Context, req dto.ShiftCreateRequest) (dto.ShiftResponse, error)
    FindAll(ctx context.Context) ([]dto.ShiftResponse, error)
    FindByID(ctx context.Context, id uuid.UUID) (dto.ShiftResponse, error)
    Update(ctx context.Context, id uuid.UUID, req dto.ShiftUpdateRequest) (dto.ShiftResponse, error)
    Delete(ctx context.Context, id uuid.UUID) error

    CreateEmployeeShift(ctx context.Context, req dto.EmployeeShiftCreateRequest) (dto.EmployeeShiftResponse, error)
    FindEmployeeShifts(ctx context.Context, employeeID uuid.UUID) ([]dto.EmployeeShiftResponse, error)
    DeleteEmployeeShift(ctx context.Context, id uuid.UUID) error
}

type shiftService struct{
    repo repository.ShiftRepository
}

func NewShiftService(r repository.ShiftRepository) ShiftService {
    return &shiftService{repo: r}
}

func (s *shiftService) Create(ctx context.Context, req dto.ShiftCreateRequest) (dto.ShiftResponse, error) {
    // parse times as stored strings (DB uses time type in entity)
    sh := &entities.Shift{
        Name: req.Name,
    }
    // For simplicity, store start/end as now placeholders; project can refine parsing
    sh.StartTime = time.Now()
    sh.EndTime = time.Now()
    sh.LateToleranceMinutes = req.LateToleranceMinutes

    if err := s.repo.Create(ctx, sh); err != nil {
        return dto.ShiftResponse{}, err
    }
    return dto.ShiftResponse{ID: sh.ID.String(), Name: sh.Name, StartTime: sh.StartTime.Format("15:04:05"), EndTime: sh.EndTime.Format("15:04:05"), LateToleranceMinutes: sh.LateToleranceMinutes}, nil
}

func (s *shiftService) FindAll(ctx context.Context) ([]dto.ShiftResponse, error) {
    list, err := s.repo.FindAll(ctx)
    if err != nil {
        return nil, err
    }
    var out []dto.ShiftResponse
    for _, sh := range list {
        out = append(out, dto.ShiftResponse{ID: sh.ID.String(), Name: sh.Name, StartTime: sh.StartTime.Format("15:04:05"), EndTime: sh.EndTime.Format("15:04:05"), LateToleranceMinutes: sh.LateToleranceMinutes})
    }
    return out, nil
}

func (s *shiftService) FindByID(ctx context.Context, id uuid.UUID) (dto.ShiftResponse, error) {
    sh, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return dto.ShiftResponse{}, err
    }
    return dto.ShiftResponse{ID: sh.ID.String(), Name: sh.Name, StartTime: sh.StartTime.Format("15:04:05"), EndTime: sh.EndTime.Format("15:04:05"), LateToleranceMinutes: sh.LateToleranceMinutes}, nil
}

func (s *shiftService) Update(ctx context.Context, id uuid.UUID, req dto.ShiftUpdateRequest) (dto.ShiftResponse, error) {
    sh, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return dto.ShiftResponse{}, err
    }
    if req.Name != nil {
        sh.Name = *req.Name
    }
    if req.LateToleranceMinutes != nil {
        sh.LateToleranceMinutes = *req.LateToleranceMinutes
    }
    if err := s.repo.Update(ctx, sh); err != nil {
        return dto.ShiftResponse{}, err
    }
    return s.FindByID(ctx, id)
}

func (s *shiftService) Delete(ctx context.Context, id uuid.UUID) error {
    return s.repo.Delete(ctx, id)
}

func (s *shiftService) CreateEmployeeShift(ctx context.Context, req dto.EmployeeShiftCreateRequest) (dto.EmployeeShiftResponse, error) {
    es := &entities.EmployeeShift{
        EmployeeID: uuid.MustParse(req.EmployeeID),
        ShiftID: uuid.MustParse(req.ShiftID),
    }
    // placeholder for effective date parsing
    es.EffectiveDate = time.Now()
    if err := s.repo.CreateEmployeeShift(ctx, es); err != nil {
        return dto.EmployeeShiftResponse{}, err
    }
    return dto.EmployeeShiftResponse{ID: es.ID.String(), EmployeeID: es.EmployeeID.String(), ShiftID: es.ShiftID.String(), EffectiveDate: es.EffectiveDate.Format("2006-01-02")}, nil
}

func (s *shiftService) FindEmployeeShifts(ctx context.Context, employeeID uuid.UUID) ([]dto.EmployeeShiftResponse, error) {
    list, err := s.repo.FindEmployeeShiftsByEmployee(ctx, employeeID)
    if err != nil {
        return nil, err
    }
    var out []dto.EmployeeShiftResponse
    for _, es := range list {
        out = append(out, dto.EmployeeShiftResponse{ID: es.ID.String(), EmployeeID: es.EmployeeID.String(), ShiftID: es.ShiftID.String(), EffectiveDate: es.EffectiveDate.Format("2006-01-02")})
    }
    return out, nil
}

func (s *shiftService) DeleteEmployeeShift(ctx context.Context, id uuid.UUID) error {
    return s.repo.DeleteEmployeeShift(ctx, id)
}
