package sqlite

import (
	"github.com/jmoiron/sqlx"
)

type ArticleSqliteRepo struct {
	db *sqlx.DB
}
