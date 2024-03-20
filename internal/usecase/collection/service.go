package collection

import (
	"newser.app/internal/infra/repository"
)

type CollectionService struct {
	collectionRepo repository.CollectionRepository
}

func NewCollectionService(collectionRepo repository.CollectionRepository) CollectionService {
	return CollectionService{
		collectionRepo: collectionRepo,
	}
}
