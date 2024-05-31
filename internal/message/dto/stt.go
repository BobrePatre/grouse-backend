package dto

type (
	SendSttNotificationRequest struct {
		UserId  int64  `json:"user_id"`
		Content string `json:"content"`
	}
)
