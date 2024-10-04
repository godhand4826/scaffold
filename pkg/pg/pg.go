package pg

import (
	"context"
	"fmt"

	_ "github.com/lib/pq" // pg driver

	"scaffold/ent"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func New(config Config) (*ent.Client, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password,
		config.Host, config.Port, config.Database,
	)

	return ent.Open("postgres", dsn)
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rErr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
