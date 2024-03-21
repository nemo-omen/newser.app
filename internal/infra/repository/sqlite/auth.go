package sqlite

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
	"newser.app/internal/domain/entity"
	"newser.app/internal/dto"
	"newser.app/shared"
)

var (
	ErrDuplicateEmail  = errors.New("email already exists")
	ErrorDuplicateName = errors.New("name already exists")
)

type AuthSqliteRepo struct {
	db *sqlx.DB
}

func NewAuthSqliteRepo(db *sqlx.DB) *AuthSqliteRepo {
	return &AuthSqliteRepo{
		db: db,
	}
}

// CreateUser creates a new user in the database.
// It also creates the default collections for the user.
func (r *AuthSqliteRepo) CreateUser(ud dto.UserDAO, collections []*entity.Collection) error {
	userQuery := `INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)`
	collectionQuery := `INSERT INTO collections (id, user_id, title, slug) VALUES (?, ?, ?, ?)`

	// Start transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Insert user
	_, err = tx.Exec(userQuery, ud.ID, ud.Name, ud.Email, ud.Password)
	if err != nil {
		fmt.Println("error: ", err)
		// check err type so we can return a more specific error
		if sqliteError, ok := err.(sqlite3.Error); ok {
			if strings.Contains(sqliteError.Error(), "UNIQUE constraint failed") {
				if strings.Contains(sqliteError.Error(), "users.email") {
					return shared.NewAppError(
						ErrDuplicateEmail,
						"That email already exists",
						"CreateUser",
						"value.Email",
					)
				}

				if strings.Contains(sqliteError.Error(), "users.name") {
					return shared.NewAppError(
						ErrorDuplicateName,
						"That name already exists",
						"CreateUser",
						"value.Name",
					)
				}
			}
		}
		return err
	}

	// Insert user collections
	for _, collection := range collections {
		_, err = tx.Exec(collectionQuery, collection.ID, collection.UserID, collection.Title.String(), collection.Slug)
		if err != nil {
			return shared.NewAppError(
				err,
				"Failed to create user collections",
				"CreateUser",
				"repository.AuthSqliteRepo",
			)
		}
	}

	tx.Commit()
	return nil
}

// FindByEmail returns a user with the given email.
func (r *AuthSqliteRepo) FindByEmail(email string) (*entity.User, error) {
	query := `SELECT id, name, email FROM users WHERE email = ? LIMIT 1`
	var u entity.User
	err := r.db.Get(&u, query, email)
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Failed to find user by email",
			"FindByEmail",
			"repository.AuthSqliteRepo",
		)
	}
	return &u, nil
}

// FindByID returns a user with the given ID.
func (r *AuthSqliteRepo) FindByID(id entity.ID) (*entity.User, error) {
	query := `SELECT * FROM users WHERE id = ? LIMIT 1`
	var u entity.User
	err := r.db.Get(&u, query, id)
	if err != nil {
		return nil, shared.NewAppError(
			err,
			"Failed to find user by ID",
			"FindByID",
			"repository.AuthSqliteRepo",
		)
	}
	return &u, nil
}

// GetHashedPasswordByEmail returns the hashed password for a user with the given email.
func (r *AuthSqliteRepo) GetHashedPasswordByEmail(email string) (string, error) {
	query := `SELECT password FROM users WHERE email = ? LIMIT 1`
	var password string
	err := r.db.Get(&password, query, email)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return "", shared.NewAppError(
				err,
				fmt.Sprintf("%s is not registered", email),
				"GetHashedPasswordByEmail",
				"value.Email",
			)
		}
	}
	return password, nil
}
