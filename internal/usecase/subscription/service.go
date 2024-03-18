package subscription

import (
	"newser.app/internal/dto"
	"newser.app/internal/infra/repository"
)

type SubscriptionService struct {
	subscriptionRepo repository.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo repository.SubscriptionRepository) SubscriptionService {
	return SubscriptionService{
		subscriptionRepo: subscriptionRepo,
	}
}

func (s *SubscriptionService) GetAllArticles(userID string) ([]*dto.ArticleDTO, error) {
	return s.subscriptionRepo.GetAllArticles(userID)
}
