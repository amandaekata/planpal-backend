package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Config holds application configuration from environment.
type Config struct {
	Port     int
	Env      string
	Database struct {
		URL string
	}
	JWT struct {
		Secret        string
		AccessExpiry  time.Duration
		RefreshExpiry time.Duration
	}
	CORS struct {
		Origins []string
	}
}

// Load reads .env and populates Config. Safe to call multiple times.
func Load() (*Config, error) {
	_ = godotenv.Load()

	c := &Config{}
	c.Port = getInt("PORT", 8080)
	c.Env = getEnv("ENV", "development")
	c.Database.URL = getEnv("DATABASE_URL", "postgres://planpal:planpal@localhost:5432/planpal?sslmode=disable")
	c.JWT.Secret = getEnv("JWT_SECRET", "change-me-in-production")
	c.JWT.AccessExpiry = getDuration("JWT_ACCESS_EXPIRY", 15*time.Minute)
	c.JWT.RefreshExpiry = getDuration("JWT_REFRESH_EXPIRY", 168*time.Hour)

	origins := getEnv("CORS_ORIGINS", "http://localhost:*")
	c.CORS.Origins = strings.Split(origins, ",")
	for i := range c.CORS.Origins {
		c.CORS.Origins[i] = strings.TrimSpace(c.CORS.Origins[i])
	}

	return c, nil
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func getInt(key string, defaultVal int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return defaultVal
}

func getDuration(key string, defaultVal time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return defaultVal
}
