package core

import (
	"log/slog"

	"newser.app/internal/repository"
	"newser.app/internal/service"
)

type App struct {
	Logger      *slog.Logger
	FeedRepo    *repository.NewsfeedSqliteRepo
	FeedService *service.API
}
