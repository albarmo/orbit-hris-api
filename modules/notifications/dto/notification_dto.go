package dto

type NotificationCreateRequest struct {
    UserID  string `json:"user_id" binding:"required,uuid"`
    Title   string `json:"title" binding:"required"`
    Message string `json:"message" binding:"required"`
}

type NotificationUpdateRequest struct {
    Title   *string `json:"title"`
    Message *string `json:"message"`
    IsRead  *bool   `json:"is_read"`
}

type NotificationResponse struct {
    ID      string `json:"id"`
    UserID  string `json:"user_id"`
    Title   string `json:"title"`
    Message string `json:"message"`
    IsRead  bool   `json:"is_read"`
}
