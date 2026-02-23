package repository

import (
	"context"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository interface {
    Create(ctx context.Context, n *entities.Notification) error
    FindAllByUser(ctx context.Context, userID uuid.UUID) ([]entities.Notification, error)
    FindByID(ctx context.Context, id uuid.UUID) (*entities.Notification, error)
    Update(ctx context.Context, n *entities.Notification) error
    Delete(ctx context.Context, id uuid.UUID) error
}

type notificationRepository struct{ db *gorm.DB }

func NewNotificationRepository(db *gorm.DB) NotificationRepository { return &notificationRepository{db: db} }

func (r *notificationRepository) Create(ctx context.Context, n *entities.Notification) error {
    return r.db.WithContext(ctx).Create(n).Error
}

func (r *notificationRepository) FindAllByUser(ctx context.Context, userID uuid.UUID) ([]entities.Notification, error) {
    var list []entities.Notification
    if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&list).Error; err != nil { return nil, err }
    return list, nil
}

func (r *notificationRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Notification, error) {
    var n entities.Notification
    if err := r.db.WithContext(ctx).First(&n, "id = ?", id).Error; err != nil { return nil, err }
    return &n, nil
}

func (r *notificationRepository) Update(ctx context.Context, n *entities.Notification) error { return r.db.WithContext(ctx).Save(n).Error }

func (r *notificationRepository) Delete(ctx context.Context, id uuid.UUID) error { return r.db.WithContext(ctx).Delete(&entities.Notification{}, "id = ?", id).Error }
