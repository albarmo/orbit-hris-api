package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/employee/dto"
	erepo "github.com/Caknoooo/go-gin-clean-starter/modules/employee/repository"
	"github.com/google/uuid"
)

type EmergencyContactService interface {
    Create(ctx context.Context, employeeID uuid.UUID, req dto.EmergencyContactCreateRequest) (dto.EmergencyContactResponse, error)
    FindByEmployee(ctx context.Context, employeeID uuid.UUID) ([]dto.EmergencyContactResponse, error)
    Delete(ctx context.Context, id uuid.UUID) error
}

type emergencyContactService struct{
    repo erepo.EmergencyContactRepository
}

func NewEmergencyContactService(r erepo.EmergencyContactRepository) EmergencyContactService {
    return &emergencyContactService{repo: r}
}

func (s *emergencyContactService) Create(ctx context.Context, employeeID uuid.UUID, req dto.EmergencyContactCreateRequest) (dto.EmergencyContactResponse, error) {
    e := &entities.EmployeeEmergencyContact{
        EmployeeID: employeeID,
        Name: req.Name,
        Relationship: req.Relationship,
        Phone: req.Phone,
    }
    if err := s.repo.Create(ctx, e); err != nil {
        return dto.EmergencyContactResponse{}, err
    }
    return dto.EmergencyContactResponse{ID: e.ID.String(), EmployeeID: e.EmployeeID.String(), Name: e.Name, Relationship: e.Relationship, Phone: e.Phone}, nil
}

func (s *emergencyContactService) FindByEmployee(ctx context.Context, employeeID uuid.UUID) ([]dto.EmergencyContactResponse, error) {
    list, err := s.repo.FindByEmployee(ctx, employeeID)
    if err != nil {
        return nil, err
    }
    var out []dto.EmergencyContactResponse
    for _, e := range list {
        out = append(out, dto.EmergencyContactResponse{ID: e.ID.String(), EmployeeID: e.EmployeeID.String(), Name: e.Name, Relationship: e.Relationship, Phone: e.Phone})
    }
    return out, nil
}

func (s *emergencyContactService) Delete(ctx context.Context, id uuid.UUID) error {
    return s.repo.Delete(ctx, id)
}
