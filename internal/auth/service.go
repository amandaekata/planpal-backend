package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Service handles JWT issue/refresh and (later) OAuth.
type Service struct {
	secret        string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
	userFinder    UserFinder
}

// UserFinder is implemented by user service to resolve user for login.
type UserFinder interface {
	FindByEmail(ctx interface{}, email string) (UserInfo, error)
}

// UserInfo is the minimal user data needed for auth.
type UserInfo struct {
	ID       string
	Email    string
	PasswordHash string
}

// NewService returns an auth service.
func NewService(secret string, accessExpiry, refreshExpiry time.Duration, userFinder UserFinder) *Service {
	return &Service{
		secret:        secret,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
		userFinder:    userFinder,
	}
}

// Claims for JWT.
type Claims struct {
	jwt.RegisteredClaims
	Email string `json:"email,omitempty"`
}

// TokenPair holds access and refresh tokens.
type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// IssueTokens creates access + refresh tokens for a user.
func (s *Service) IssueTokens(userID, email string) (*TokenPair, error) {
	now := time.Now()
	accessExp := now.Add(s.accessExpiry)
	refreshExp := now.Add(s.refreshExpiry)

	accessClaims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(accessExp),
			ID:        uuid.New().String(),
		},
		Email: email,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	refreshClaims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(refreshExp),
			ID:        uuid.New().String(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
		ExpiresAt:    accessExp,
	}, nil
}
