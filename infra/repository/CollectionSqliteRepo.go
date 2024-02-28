package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type CollectionSqliteRepo struct {
	DB *sqlx.DB
}

func (r CollectionSqliteRepo) Migrate() error {
	fmt.Println("Migrating collections table...")
	q := `
	CREATE TABLE IF NOT EXISTS collections(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		slug TEXT NOT NULL,
		user_id INT NOT NULL,
		CONSTRAINT fk_users
			FOREIGN KEY (user_id)
				REFERENCES users (id)
	)
	`
	_, err := r.DB.Exec(q)
	if err != nil {
		fmt.Println("error migrating collections: ", err.Error())
		return err
	} else {
		fmt.Println("completed migrating collections")
	}
	return err
}

func NewCollectionSqliteRepo(db *sqlx.DB) CollectionSqliteRepo {
	return CollectionSqliteRepo{DB: db}
}

func (r CollectionSqliteRepo) Get(id int64) (model.Collection, error) {
	return model.Collection{}, nil
}

func (r CollectionSqliteRepo) Create(c model.Collection) (model.Collection, error) {
	q := `
	INSERT INTO collections(title, slug, user_id)
		VALUES(?, ?, ?);
	`
	res, err := r.DB.Exec(q, c.Title, c.Slug, c.UserId)

	if err != nil {
		return model.Collection{}, err
	}
	id, _ := res.LastInsertId()
	c.Id = id
	return c, nil
}

func (r CollectionSqliteRepo) All(userId int64) ([]model.Collection, error) {
	ss := []model.Collection{}
	err := r.DB.Select(&ss, "SELECT * FROM collections WHERE user_id=?", userId)
	if err != nil {
		return ss, err
	}
	return ss, nil
}

func (r CollectionSqliteRepo) Update(s model.Collection) (model.Collection, error) {
	return model.Collection{}, nil
}

func (r CollectionSqliteRepo) Delete(id int64) error {
	return nil
}

func (r CollectionSqliteRepo) FindBySlug(slug string) (model.Collection, error) {
	return model.Collection{}, nil
}
