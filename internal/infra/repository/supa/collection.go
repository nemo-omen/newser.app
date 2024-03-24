package supa

import (
	"fmt"

	supa "github.com/nedpals/supabase-go"
	"newser.app/internal/dto"
	"newser.app/shared"
)

type CollectionSupaRepo struct {
	db *supa.Client
}

func NewCollectionSupaRepo(client *supa.Client) *CollectionSupaRepo {
	return &CollectionSupaRepo{client}
}

func (r CollectionSupaRepo) Create(collection *dto.CollectionDTO) error {
	var results []*dto.CollectionDTO
	err := r.db.DB.From("collections").Insert(collection).Execute(&results)
	if err != nil {
		// return err
		return shared.NewAppError(
			err,
			"error creating collection",
			"CollectionSupaRepo.Create",
			"SupabaseError",
		)
	}

	return nil
}

func (r CollectionSupaRepo) Delete(collectionID string) error {
	res := []*dto.CollectionDTO{}
	err := r.db.DB.From("collections").Delete().Eq("id", collectionID).Execute(&res)
	if err != nil {
		// return err
		return shared.NewAppError(
			err,
			"error deleting collection",
			"CollectionSupaRepo.Delete",
			"SupabaseError",
		)
	}
	return nil
}

func (r CollectionSupaRepo) Get(collectionID string) (*dto.CollectionDTO, error) {
	res := []*dto.CollectionDTO{}
	err := r.db.DB.From("collections").Select().Eq("id", collectionID).Execute(&res)
	if err != nil {
		// return nil, err
		return nil, shared.NewAppError(
			err,
			"error getting collection",
			"CollectionSupaRepo.Get",
			"SupabaseError",
		)
	}
	if len(res) == 0 {
		return nil, nil
	}

	return res[0], nil
}

func (r CollectionSupaRepo) GetByTitle(title, userId string) (*dto.CollectionDTO, error) {
	fmt.Println("GetByTitle not implemented yet.")
	return nil, nil
}

func (r CollectionSupaRepo) GetBySlug(slug, userId string) (*dto.CollectionDTO, error) {
	fmt.Println("GetBySlug not implemented yet.")
	return nil, nil
}

func (r CollectionSupaRepo) All(userID string) ([]*dto.CollectionDTO, error) {
	fmt.Println("All not implemented yet.")
	return nil, nil
}

func (r CollectionSupaRepo) AddArticle(collectionID, articleID string) error {
	fmt.Println("AddArticle not implemented yet.")
	return nil
}

func (r CollectionSupaRepo) RemoveArticle(collectionID, articleID string) error {
	fmt.Println("RemoveArticle not implemented yet.")
	return nil
}

func (r CollectionSupaRepo) AddAndRemoveArticle(addCollectionID, removeCollectionID, articleID string) error {
	fmt.Println("AddAndRemoveArticle not implemented yet.")
	return nil
}

func (r CollectionSupaRepo) AddNewsfeed(collectionID, newsfeedID string) error {
	fmt.Println("AddNewsfeed not implemented yet.")
	return nil
}

func (r CollectionSupaRepo) RemoveNewsfeed(collectionID, newsfeedID string) error {
	fmt.Println("RemoveNewsfeed not implemented yet.")
	return nil
}

func (r CollectionSupaRepo) GetCollectionArticles(collectionID, userID string) ([]*dto.ArticleDTO, error) {
	fmt.Println("GetCollectionArticles not implemented yet.")
	return nil, nil
}

func (r CollectionSupaRepo) GetCollectionNewsfeeds(collectionID string) ([]*dto.NewsfeedDTO, error) {
	fmt.Println("GetCollectionNewsfeeds not implemented yet.")
	return nil, nil
}
