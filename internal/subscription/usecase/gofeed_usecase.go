package usecase

import (
	"newser.app/internal/subscription"
)

type GoFeedUsecase struct {
	repo subscription.SearchRepository
}
