package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PasswordResetRepository interface {
	Create(ctx context.Context, r *entities.PasswordReset) error
	FindByToken(ctx context.Context, token uuid.UUID) (*entities.PasswordReset, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type passwordResetRepository struct {
	db *gorm.DB
}

func NewPasswordResetRepository(db *gorm.DB) PasswordResetRepository {
	return &passwordResetRepository{db: db}
}

func (r *passwordResetRepository) Create(ctx context.Context, rr *entities.PasswordReset) error {
	return r.db.WithContext(ctx).Create(rr).Error
}

func (r *passwordResetRepository) FindByToken(ctx context.Context, token uuid.UUID) (*entities.PasswordReset, error) {
	var pr entities.PasswordReset
	if err := r.db.WithContext(ctx).First(&pr, "reset_token = ?", token).Error; err != nil {
		return nil, err
	}
	return &pr, nil
}

func (r *passwordResetRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.PasswordReset{}, "id = ?", id).Error
}
