package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	"newser.app/infra/dto"
	"newser.app/model"
)

var (
	ErrDuplicateEmail = errors.New("repository: duplicate email")
	ErrNotExists      = errors.New("repository: user not found")
)

type UserSqliteRepo struct {
	db *sqlx.DB
}

func NewUserSqliteRepo(db *sqlx.DB) *UserSqliteRepo {
	return &UserSqliteRepo{db: db}
}

func (r *UserSqliteRepo) Migrate() error {
	fmt.Println("Starting users table migration...")
	q := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err := r.db.Exec(q)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("users migration complete.")
	}
	return nil
}

func (r *UserSqliteRepo) Get(id int64) (*model.User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE id=?", id)
	var u model.User
	if err := row.Scan(&u.Id, &u.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserSqliteRepo) FindByEmail(email string) (*model.User, error) {
	u := model.User{}
	err := r.db.Get(&u, "SELECT id, email FROM users WHERE email=?", email)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserSqliteRepo) GetHashedPasswordByEmail(email string) (string, error) {
	var hashed string
	err := r.db.Get(&hashed, "SELECT password FROM users WHERE email=?", email)
	if err != nil {
		fmt.Println("err: ", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNotExists
		}
		return "", nil
	}
	return hashed, nil
}

func (r *UserSqliteRepo) Create(udto *dto.UserDTO) (*model.User, error) {
	var u model.User
	q := `
	INSERT INTO users(email, password)
		VALUES(?, ?)
	`
	res, err := r.db.Exec(q, udto.Email, udto.HashedPassword)
	if err != nil {
		var sqliteError *sqlite3.Error
		if errors.As(err, &sqliteError) {
			if sqliteError.Code == 2067 {
				return nil, ErrDuplicateEmail
			}
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	u.Email = udto.Email
	u.Id = id
	return &u, nil
}

func (r *UserSqliteRepo) All() []*model.User {
	return nil
}

func (r *UserSqliteRepo) Update(udto *dto.UserDTO) (*model.User, error) {
	return nil, nil
}

func (r *UserSqliteRepo) Delete(id int64) error {
	return nil
}
