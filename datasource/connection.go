package datasource

import (
	"context"

	"github.com/vingarcia/ksql"
	ksqlite "github.com/vingarcia/ksql/adapters/modernc-ksqlite"
	"newser.app/config"
)

func NewDatabase(cfg config.DatabaseConfig) (*ksql.DB, error) {
	ctx := context.Background()
	db, err := ksqlite.New(ctx, cfg.DSN, ksql.Config{})
	if err != nil {
		return nil, err
	}
	return &db, nil
}
