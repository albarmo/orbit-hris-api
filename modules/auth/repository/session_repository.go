package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type SessionRepository interface {
	Create(ctx context.Context, tx any, token string, userID uuid.UUID, expiresAt time.Time) error
	GetUserIDByAccessToken(ctx context.Context, token string) (uuid.UUID, error)
	DeleteByToken(ctx context.Context, token string) error
	DeleteByUserID(ctx context.Context, userID string) error
}

type sessionRepo struct {
	redisClient *redis.Client
}

func NewSessionRepository(redisClient *redis.Client) SessionRepository {
	return &sessionRepo{redisClient: redisClient}
}

type redisSessionValue struct {
	UserID string `json:"user_id"`
}

func (r *sessionRepo) Create(ctx context.Context, tx any, token string, userID uuid.UUID, expiresAt time.Time) error {
	key := fmt.Sprintf("access:%s", token)
	val := redisSessionValue{UserID: userID.String()}
	b, _ := json.Marshal(val)

	ttl := time.Until(expiresAt)
	if ttl <= 0 {
		ttl = time.Minute * 15
	}

	return r.redisClient.Set(ctx, key, b, ttl).Err()
}

func (r *sessionRepo) GetUserIDByAccessToken(ctx context.Context, token string) (uuid.UUID, error) {
	key := fmt.Sprintf("access:%s", token)
	res, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return uuid.Nil, err
	}
	var rv redisSessionValue
	if err := json.Unmarshal([]byte(res), &rv); err != nil {
		return uuid.Nil, err
	}
	uid, err := uuid.Parse(rv.UserID)
	if err != nil {
		return uuid.Nil, err
	}
	return uid, nil
}

func (r *sessionRepo) DeleteByToken(ctx context.Context, token string) error {
	key := fmt.Sprintf("access:%s", token)
	return r.redisClient.Del(ctx, key).Err()
}

func (r *sessionRepo) DeleteByUserID(ctx context.Context, userID string) error {
	// iterate keys pattern access:* and check user id (not ideal for large scale)
	// better to maintain a set per user; implement best-effort here.
	pattern := "access:*"
	iter := r.redisClient.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		k := iter.Val()
		res, err := r.redisClient.Get(ctx, k).Result()
		if err != nil {
			continue
		}
		var rv redisSessionValue
		if err := json.Unmarshal([]byte(res), &rv); err != nil {
			continue
		}
		if rv.UserID == userID {
			r.redisClient.Del(ctx, k)
		}
	}
	return iter.Err()
}

// NewNoopSessionRepository returns a no-op repo for environments without Redis
func NewNoopSessionRepository() SessionRepository {
	return &noopSessionRepo{}
}

type noopSessionRepo struct{}

func (n *noopSessionRepo) Create(ctx context.Context, tx any, token string, userID uuid.UUID, expiresAt time.Time) error {
	return nil
}
func (n *noopSessionRepo) GetUserIDByAccessToken(ctx context.Context, token string) (uuid.UUID, error) {
	return uuid.Nil, fmt.Errorf("not available")
}
func (n *noopSessionRepo) DeleteByToken(ctx context.Context, token string) error   { return nil }
func (n *noopSessionRepo) DeleteByUserID(ctx context.Context, userID string) error { return nil }
