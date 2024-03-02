package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"newser.app/model"
)

type PersonRepository interface {
	Create(name, email string) (*model.Person, error)
	Get(id int64) (*model.Person, error)
	GetByName(name string) (*model.Person, error)
	GetByEmail(email string) (*model.Person, error)
	Migrate() error
}

type PersonSqliteRepository struct {
	db *sqlx.DB
}

func NewPersonSqliteRepository(db *sqlx.DB) *PersonSqliteRepository {
	return &PersonSqliteRepository{
		db: db,
	}
}

func (r PersonSqliteRepository) Create(name, email string) (*model.Person, error) {
	q := `
	INSERT INTO people(name, email)
		VALUES(?, ?);
	`
	p := model.Person{Name: name, Email: email}
	res, err := r.db.Exec(q, &p)
	if err != nil {
		return nil, ErrInsertError
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrInsertError
	}
	p.ID = id
	return &p, nil
}

func (r *PersonSqliteRepository) Get(id int64) (*model.Person, error) {
	p := model.Person{}
	err := r.db.Get(&p, "SELECT * FROM people WHERE id=?", id)
	if err != nil {
		return nil, ErrNotExists
	}
	return &p, nil
}

func (r *PersonSqliteRepository) GetByName(name string) (*model.Person, error) {
	p := model.Person{}
	err := r.db.Get(&p, "SELECT * FROM people WHERE name=?", name)
	if err != nil {
		return nil, ErrNotExists
	}
	return &p, nil
}

func (r *PersonSqliteRepository) GetByEmail(email string) (*model.Person, error) {
	p := model.Person{}
	err := r.db.Get(&p, "SELECT * FROM people WHERE email=?", email)
	if err != nil {
		return nil, ErrNotExists
	}
	return &p, nil
}

func (r *PersonSqliteRepository) Migrate() error {
	fmt.Println("Starting person migrations")
	q1 := `
	CREATE TABLE IF NOT EXISTS people(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE ON CONFLICT IGNORE
	);
	`
	_, err := r.db.Exec(q1)
	if err != nil {
		fmt.Println("error migrating people table: ", err.Error())
		return ErrMigrationError
	}

	q2 := `
	CREATE TABLE IF NOT EXISTS article_people(
		person_id INTEGER NOT NULL,
		article_id INTEGER NOT NULL,
		PRIMARY KEY(person_id, article_id),
		CONSTRAINT fk_articles
			FOREIGN KEY(article_id)
			REFERENCES articles(id),
		CONSTRAINT fk_people
			FOREIGN KEY(PERSON_ID)
			REFERENCES people(id)
	);
	`
	_, err = r.db.Exec(q2)
	if err != nil {
		fmt.Println("error migrating article_people table: ", err.Error())
		return ErrMigrationError
	}
	q3 := `
	CREATE TABLE IF NOT EXISTS newsfeed_people(
		person_id INTEGER NOT NULL,
		newsfeed_id INTEGER NOT NULL,
		PRIMARY KEY(person_id, newsfeed_id),
		CONSTRAINT fk_newsfeeds
			FOREIGN KEY(newsfeed_id)
			REFERENCES newsfeeds(id),
		CONSTRAINT fk_people
			FOREIGN KEY(PERSON_ID)
			REFERENCES people(id)
	);
	`
	_, err = r.db.Exec(q3)
	if err != nil {
		fmt.Println("error migrating newsfeed_people table: ", err.Error())
		return ErrMigrationError
	}
	fmt.Println("Completed migrating person tables")
	return nil
}
