package pg

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/lib/pq" // pg driver

	"scaffold/ent"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Debug    bool
}

func New(config Config, logger *zap.Logger) (*ent.Client, error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password,
		config.Host, config.Port, config.Database,
	)

	client, err := ent.Open("postgres", dsn, ent.Log(ZapToEntLogger(logger)))
	if err != nil {
		return nil, err
	}

	if config.Debug {
		client = client.Debug()
	}

	return client, nil
}

func ZapToEntLogger(logger *zap.Logger) func(...any) {
	return logger.Sugar().Debug
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
