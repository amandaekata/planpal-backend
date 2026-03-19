package streak

// Service handles streak engine logic (daily/weekly completion, streak counts).
type Service struct {
	repo Repository
}

// Repository interface for streak persistence.
type Repository interface {
	GetByUser(ctx interface{}, userID string) (*Streak, error)
	Upsert(ctx interface{}, s *Streak) error
	RecordCompletion(ctx interface{}, userID, goalID, date string) error
}

// Streak holds current streak state for a user.
type Streak struct {
	UserID       string `json:"user_id"`
	Current      int    `json:"current"`
	Longest      int    `json:"longest"`
	LastDoneDate string `json:"last_done_date,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// NewService returns a streak service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
