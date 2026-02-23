package service

import (
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/v4/messaging"
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/Caknoooo/go-gin-clean-starter/modules/notifications/dto"
	"github.com/Caknoooo/go-gin-clean-starter/modules/notifications/repository"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/google/uuid"
	"github.com/samber/do"
)

type NotificationService interface {
	Create(ctx context.Context, req dto.NotificationCreateRequest) (dto.NotificationResponse, error)
	GetByUser(ctx context.Context, userID uuid.UUID) ([]dto.NotificationResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.NotificationResponse, error)
	Update(ctx context.Context, id uuid.UUID, req dto.NotificationUpdateRequest) (dto.NotificationResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type notificationService struct {
	repo repository.NotificationRepository
	fcm  *messaging.Client
}

func NewNotificationService(r repository.NotificationRepository, injector *do.Injector) NotificationService {
	// try to obtain fcm client; if not available (e.g., during migrations),
	// proceed with nil client so service still constructs.
	var fcm *messaging.Client
	if injector != nil {
		if c, err := do.InvokeNamed[*messaging.Client](injector, constants.FCMClient); err == nil {
			fcm = c
		}
	}
	return &notificationService{repo: r, fcm: fcm}
}

func (s *notificationService) Create(ctx context.Context, req dto.NotificationCreateRequest) (dto.NotificationResponse, error) {
	n := &entities.Notification{UserID: uuid.MustParse(req.UserID), Title: req.Title, Message: req.Message, IsRead: false, CreatedAt: time.Now()}
	if err := s.repo.Create(ctx, n); err != nil {
		return dto.NotificationResponse{}, err
	}

	// send FCM to topic for the user: topic name `user_<userID>`
	if s.fcm != nil {
		topic := fmt.Sprintf("user_%s", n.UserID.String())
		msg := &messaging.Message{
			Notification: &messaging.Notification{Title: n.Title, Body: n.Message},
			Topic:        topic,
		}
		// best-effort: log error but don't fail the overall create
		if _, err := s.fcm.Send(ctx, msg); err != nil {
			// ignore send error for now
		}
	}

	return dto.NotificationResponse{ID: n.ID.String(), UserID: n.UserID.String(), Title: n.Title, Message: n.Message, IsRead: n.IsRead}, nil
}

func (s *notificationService) GetByUser(ctx context.Context, userID uuid.UUID) ([]dto.NotificationResponse, error) {
	list, err := s.repo.FindAllByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	var out []dto.NotificationResponse
	for _, n := range list {
		out = append(out, dto.NotificationResponse{ID: n.ID.String(), UserID: n.UserID.String(), Title: n.Title, Message: n.Message, IsRead: n.IsRead})
	}
	return out, nil
}

func (s *notificationService) GetByID(ctx context.Context, id uuid.UUID) (dto.NotificationResponse, error) {
	n, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.NotificationResponse{}, err
	}
	return dto.NotificationResponse{ID: n.ID.String(), UserID: n.UserID.String(), Title: n.Title, Message: n.Message, IsRead: n.IsRead}, nil
}

func (s *notificationService) Update(ctx context.Context, id uuid.UUID, req dto.NotificationUpdateRequest) (dto.NotificationResponse, error) {
	n, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.NotificationResponse{}, err
	}
	if req.Title != nil {
		n.Title = *req.Title
	}
	if req.Message != nil {
		n.Message = *req.Message
	}
	if req.IsRead != nil {
		n.IsRead = *req.IsRead
	}
	if err := s.repo.Update(ctx, n); err != nil {
		return dto.NotificationResponse{}, err
	}
	return s.GetByID(ctx, id)
}

func (s *notificationService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
