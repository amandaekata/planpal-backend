package goal

// Service handles goals CRUD. Repository is injected when DB is wired.
type Service struct {
	repo Repository
}

// Repository interface for goal persistence.
type Repository interface {
	GetByID(ctx interface{}, id, userID string) (*Goal, error)
	ListByUser(ctx interface{}, userID string) ([]*Goal, error)
	Create(ctx interface{}, g *Goal) error
	Update(ctx interface{}, g *Goal) error
	Delete(ctx interface{}, id, userID string) error
}

// Goal represents a weekly/daily goal.
type Goal struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"` // daily, weekly
	TargetCount int    `json:"target_count"`
	Completed   int    `json:"completed"`
	WeekStart   string `json:"week_start,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// NewService returns a goal service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
