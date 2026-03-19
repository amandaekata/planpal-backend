package user

import "github.com/amandaekata/planpal-backend/internal/auth"

// Service handles user CRUD and profile. Repository is injected when DB is wired.
type Service struct {
	repo Repository
}

// Repository interface for user persistence.
type Repository interface {
	GetByID(ctx interface{}, id string) (*User, error)
	GetByEmail(ctx interface{}, email string) (*User, error)
	Create(ctx interface{}, u *User) error
	Update(ctx interface{}, u *User) error
}

// User represents a PlanPal user.
type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	DisplayName  string `json:"display_name,omitempty"`
	PasswordHash string `json:"-"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// NewService returns a user service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetByID returns a user by ID.
func (s *Service) GetByID(ctx interface{}, id string) (*User, error) {
	if s.repo == nil {
		return nil, nil
	}
	return s.repo.GetByID(ctx, id)
}

// FindByEmail implements auth.UserFinder for login.
func (s *Service) FindByEmail(ctx interface{}, email string) (auth.UserInfo, error) {
	if s.repo == nil {
		return auth.UserInfo{}, nil
	}
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil || u == nil {
		return auth.UserInfo{}, err
	}
	return auth.UserInfo{ID: u.ID, Email: u.Email, PasswordHash: u.PasswordHash}, nil
}
