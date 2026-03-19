package notification

import "github.com/amandaekata/planpal-backend/internal/ws"

// Service handles push and in-app notifications; can broadcast via WebSocket hub.
type Service struct {
	repo Repository
	hub  *ws.Hub
}

// Repository interface for notification persistence.
type Repository interface {
	Create(ctx interface{}, n *Notification) error
	ListByUser(ctx interface{}, userID string, limit int) ([]*Notification, error)
	MarkRead(ctx interface{}, id, userID string) error
}

// Notification represents a single notification.
type Notification struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Type      string `json:"type"`
	Read      bool   `json:"read"`
	CreatedAt string `json:"created_at,omitempty"`
}

// NewService returns a notification service.
func NewService(repo Repository, hub *ws.Hub) *Service {
	return &Service{repo: repo, hub: hub}
}
