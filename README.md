# PlanPal Backend

Backend API for **PlanPal** — a weekly planner that helps you crush your daily and weekly goals while keeping things fun through gamified streaks.

Designed to connect with the [PlanPal Flutter app](https://github.com/amandaekata/planpal).

## Stack

- **Go 1.22+**
- **PostgreSQL** (with migrations via [golang-migrate](https://github.com/golang-migrate/migrate))
- **JWT** for auth, optional OAuth
- **WebSockets** for real-time updates (streaks, notifications)
- **Docker** for local dev

## Project structure

```
planpal-backend/
├── cmd/api/              # Entry point (main.go)
├── internal/
│   ├── auth/             # JWT, OAuth, refresh tokens
│   ├── user/             # User service + repository
│   ├── goal/             # Goals CRUD
│   ├── streak/           # Streak engine logic
│   ├── reward/           # XP, badges, leaderboard
│   ├── notification/     # Push + in-app alerts
│   └── ws/               # WebSocket hub
├── db/
│   ├── migrations/       # SQL migrations
│   └── queries/          # Raw SQL / SQLC queries
├── config/               # Env config loading
├── middleware/           # Auth, logging, rate limiting
└── docker-compose.yml
```

## Quick start

1. **Copy env and start DB:**

   ```bash
   cp .env.example .env
   docker compose up -d
   ```

2. **Run migrations:**

   ```bash
   make migrate-up
   # or: migrate -path db/migrations -database "postgres://user:pass@localhost:5432/planpal?sslmode=disable" up
   ```

3. **Run the API:**

   ```bash
   go run ./cmd/api
   ```

API will be at `http://localhost:8080` (or the port in `.env`).

## Environment

See `.env.example` for required variables (`DATABASE_URL`, `JWT_SECRET`, etc.).

## License

MIT
