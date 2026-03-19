package reward

// Service handles XP, badges, and leaderboard.
type Service struct {
	repo Repository
}

// Repository interface for reward persistence.
type Repository interface {
	GetXP(ctx interface{}, userID string) (int, error)
	AddXP(ctx interface{}, userID string, amount int, reason string) error
	ListBadges(ctx interface{}, userID string) ([]*Badge, error)
	Leaderboard(ctx interface{}, limit int) ([]*LeaderboardEntry, error)
}

// Badge represents an earned badge.
type Badge struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon,omitempty"`
	EarnedAt    string `json:"earned_at,omitempty"`
}

// LeaderboardEntry is one row on the leaderboard.
type LeaderboardEntry struct {
	UserID    string `json:"user_id"`
	DisplayName string `json:"display_name"`
	XP        int    `json:"xp"`
	Rank      int    `json:"rank"`
}

// NewService returns a reward service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
