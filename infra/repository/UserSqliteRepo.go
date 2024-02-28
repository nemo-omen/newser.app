package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"newser.app/infra/dto"
	"newser.app/model"
)

var (
	ErrDuplicateEmail = errors.New("repository: duplicate email")
	ErrNotExists      = errors.New("repository: user not found")
)

type UserSqliteRepo struct {
	DB *sql.DB
}

func NewUserSqliteRepo(db *sql.DB) UserSqliteRepo {
	return UserSqliteRepo{DB: db}
}

func (r UserSqliteRepo) Migrate() error {
	fmt.Println("Starting users table migration...")
	q := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`

	_, err := r.DB.Exec(q)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("users migration complete.")
	}
	return err
}

func (r UserSqliteRepo) Get(id int64) (model.User, error) {
	row := r.DB.QueryRow("SELECT * FROM users WHERE id=?", id)
	var u model.User
	if err := row.Scan(&u.Id, &u.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, ErrNotExists
		}
		return u, err
	}
	return u, nil
}

func (r UserSqliteRepo) FindByEmail(email string) (model.User, error) {
	row := r.DB.QueryRow("SELECT * FROM users WHERE email=?", email)
	var u model.User
	if err := row.Scan(&u.Id, &u.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, ErrNotExists
		}
		return u, err
	}
	return u, nil
}

func (r UserSqliteRepo) GetHashedPasswordByEmail(email string) (string, error) {
	row := r.DB.QueryRow("SELECT * FROM users WHERE email=?", email)
	var u dto.UserDTO
	if err := row.Scan(&u.Id, &u.Email, &u.HashedPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNotExists
		}
		return "", err
	}
	return u.HashedPassword, nil
}

func (r UserSqliteRepo) Create(udto dto.UserDTO) (model.User, error) {
	var u model.User
	q := `
	INSERT INTO users(email, password)
		VALUES(?, ?)
	`
	res, err := r.DB.Exec(q, udto.Email, udto.HashedPassword)
	if err != nil {
		var sqliteError *sqlite3.Error
		if errors.As(err, &sqliteError) {
			if sqliteError.Code == 2067 {
				return u, ErrDuplicateEmail
			}
		}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.Email = udto.Email
	u.Id = id
	return u, nil
}

func (r UserSqliteRepo) All() []model.User {
	return []model.User{}
}

func (r UserSqliteRepo) Update(udto dto.UserDTO) (model.User, error) {
	return model.User{}, nil
}

func (r UserSqliteRepo) Delete(id int64) error {
	return nil
}
