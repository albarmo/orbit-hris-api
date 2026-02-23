package dto

import "time"

type PasswordResetCreateRequest struct {
    UserID    string    `json:"user_id" binding:"required,uuid"`
    ExpiresAt time.Time `json:"expires_at" binding:"required"`
}

type PasswordResetResponse struct {
    ID         string     `json:"id"`
    UserID     string     `json:"user_id"`
    ResetToken string     `json:"reset_token"`
    ExpiresAt  time.Time  `json:"expires_at"`
    UsedAt     *time.Time `json:"used_at"`
    CreatedAt  time.Time  `json:"created_at"`
}
