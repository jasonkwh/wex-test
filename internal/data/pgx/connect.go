package pgx

import (
	"context"
	"fmt"

	"github.com/jasonkwh/wex-test/internal/config"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg config.DatabaseConfig, zl *zap.Logger) (*pgxpool.Pool, error) {
	pcfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s", cfg.Host, cfg.Port, cfg.Database, cfg.User, cfg.Password))
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(ctx, pcfg)
}
