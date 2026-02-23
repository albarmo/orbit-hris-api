package service

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/dto"
	authrepo "github.com/Caknoooo/go-gin-clean-starter/modules/auth/repository"
	"github.com/google/uuid"
)

type PasswordResetService interface {
	Create(ctx context.Context, req dto.PasswordResetCreateRequest) (dto.PasswordResetResponse, error)
	FindByToken(ctx context.Context, token uuid.UUID) (dto.PasswordResetResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type passwordResetService struct {
	repo authrepo.PasswordResetRepository
}

func NewPasswordResetService(r authrepo.PasswordResetRepository) PasswordResetService {
	return &passwordResetService{repo: r}
}

func (s *passwordResetService) Create(ctx context.Context, req dto.PasswordResetCreateRequest) (dto.PasswordResetResponse, error) {
	token := uuid.New()
	pr := &entities.PasswordReset{
		UserID:     uuid.MustParse(req.UserID),
		ResetToken: token,
		ExpiresAt:  req.ExpiresAt,
	}
	if err := s.repo.Create(ctx, pr); err != nil {
		return dto.PasswordResetResponse{}, err
	}
	return dto.PasswordResetResponse{
		ID:         pr.ID.String(),
		UserID:     pr.UserID.String(),
		ResetToken: pr.ResetToken.String(),
		ExpiresAt:  pr.ExpiresAt,
		CreatedAt:  pr.CreatedAt,
	}, nil
}

func (s *passwordResetService) FindByToken(ctx context.Context, token uuid.UUID) (dto.PasswordResetResponse, error) {
	pr, err := s.repo.FindByToken(ctx, token)
	if err != nil {
		return dto.PasswordResetResponse{}, err
	}
	return dto.PasswordResetResponse{
		ID:         pr.ID.String(),
		UserID:     pr.UserID.String(),
		ResetToken: pr.ResetToken.String(),
		ExpiresAt:  pr.ExpiresAt,
		UsedAt:     pr.UsedAt,
		CreatedAt:  pr.CreatedAt,
	}, nil
}

func (s *passwordResetService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
