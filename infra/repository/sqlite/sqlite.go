package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (sqlx., error)
}

func transaction(ctx context.Context, tx sqlx.Tx, f func() error) error {
	if err := f(); err != nil {
		_ = tx.Rollback()

		return fmt.Errorf("transaction fn error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit error: %w", err)
	}

	return nil
}
