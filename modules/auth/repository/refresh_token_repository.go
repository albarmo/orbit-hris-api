package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	// Create stores a refresh token for a user with an expiry.
	Create(ctx context.Context, tx *gorm.DB, token string, userID uuid.UUID, expiresAt time.Time) error
	// FindByToken returns the user ID associated with the refresh token.
	FindByToken(ctx context.Context, tx *gorm.DB, token string) (uuid.UUID, error)
	DeleteByUserID(ctx context.Context, tx *gorm.DB, userID string) error
	DeleteByToken(ctx context.Context, tx *gorm.DB, token string) error
	DeleteExpired(ctx context.Context, tx *gorm.DB) error
}

type refreshTokenRepository struct {
	redisClient *redis.Client
}

func NewRefreshTokenRepository(redisClient *redis.Client) RefreshTokenRepository {
	return &refreshTokenRepository{
		redisClient: redisClient,
	}
}

type redisRefreshValue struct {
	UserID string `json:"user_id"`
}

func (r *refreshTokenRepository) Create(
	ctx context.Context,
	tx *gorm.DB,
	token string,
	userID uuid.UUID,
	expiresAt time.Time,
) error {
	if token == "" {
		return errors.New("empty token")
	}

	key := fmt.Sprintf("refresh:%s", token)
	val := redisRefreshValue{UserID: userID.String()}
	b, _ := json.Marshal(val)

	ttl := time.Until(expiresAt)
	if ttl <= 0 {
		ttl = time.Minute * 5
	}

	if err := r.redisClient.Set(ctx, key, b, ttl).Err(); err != nil {
		return err
	}

	// add token to user's set
	userSet := fmt.Sprintf("user_refreshs:%s", userID.String())
	if err := r.redisClient.SAdd(ctx, userSet, token).Err(); err != nil {
		// best-effort: remove key if set failed
		r.redisClient.Del(ctx, key)
		return err
	}

	return nil
}

func (r *refreshTokenRepository) FindByToken(ctx context.Context, tx *gorm.DB, token string) (uuid.UUID, error) {
	key := fmt.Sprintf("refresh:%s", token)
	res, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return uuid.Nil, gorm.ErrRecordNotFound
		}
		return uuid.Nil, err
	}

	var rv redisRefreshValue
	if err := json.Unmarshal([]byte(res), &rv); err != nil {
		return uuid.Nil, err
	}

	uid, err := uuid.Parse(rv.UserID)
	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}

func (r *refreshTokenRepository) DeleteByUserID(ctx context.Context, tx *gorm.DB, userID string) error {
	userSet := fmt.Sprintf("user_refreshs:%s", userID)
	tokens, err := r.redisClient.SMembers(ctx, userSet).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	// delete each token key
	for _, t := range tokens {
		key := fmt.Sprintf("refresh:%s", t)
		r.redisClient.Del(ctx, key)
	}

	// remove the set
	r.redisClient.Del(ctx, userSet)
	return nil
}

func (r *refreshTokenRepository) DeleteByToken(ctx context.Context, tx *gorm.DB, token string) error {
	key := fmt.Sprintf("refresh:%s", token)
	res, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	var rv redisRefreshValue
	if err := json.Unmarshal([]byte(res), &rv); err == nil {
		userSet := fmt.Sprintf("user_refreshs:%s", rv.UserID)
		r.redisClient.SRem(ctx, userSet, token)
	}

	r.redisClient.Del(ctx, key)
	return nil
}

func (r *refreshTokenRepository) DeleteExpired(ctx context.Context, tx *gorm.DB) error {
	// Redis uses TTL; expired keys are removed automatically. No-op.
	return nil
}

// NewNoopRefreshTokenRepository returns a repository suitable for environments
// where Redis is not available (e.g., running DB migrations). Methods are
// best-effort no-ops to allow services to be constructed.
func NewNoopRefreshTokenRepository() RefreshTokenRepository {
	return &noopRefreshRepo{}
}

type noopRefreshRepo struct{}

func (n *noopRefreshRepo) Create(ctx context.Context, tx *gorm.DB, token string, userID uuid.UUID, expiresAt time.Time) error {
	return nil
}

func (n *noopRefreshRepo) FindByToken(ctx context.Context, tx *gorm.DB, token string) (uuid.UUID, error) {
	return uuid.Nil, gorm.ErrRecordNotFound
}

func (n *noopRefreshRepo) DeleteByUserID(ctx context.Context, tx *gorm.DB, userID string) error {
	return nil
}

func (n *noopRefreshRepo) DeleteByToken(ctx context.Context, tx *gorm.DB, token string) error {
	return nil
}

func (n *noopRefreshRepo) DeleteExpired(ctx context.Context, tx *gorm.DB) error {
	return nil
}
