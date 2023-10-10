package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jasonkwh/wex-test/internal/config"
	"github.com/jasonkwh/wex-test/internal/data/model"
	"go.uber.org/zap"
)

type dbPurchaseRepository struct {
	pool *pgxpool.Pool
	zl   *zap.Logger
}

func CreatePurchaseRepository(dbcfg config.DatabaseConfig, zl *zap.Logger) (PurchaseRepository, error) {
	ctx := context.Background()

	pool, err := Connect(ctx, dbcfg, zl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to wex db: %v", err)
	}

	return &dbPurchaseRepository{
		pool: pool,
		zl:   zl,
	}, nil
}

func (r *dbPurchaseRepository) SavePurchase(ctx context.Context, p *model.Transaction) (string, error) {
	r.zl.Info("repo.SavePurchase: started")

	// start db tx
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return "", err
	}

	// the unique identifier
	id := p.Hash()

	_, err = tx.Exec(ctx, "select wex.save_purchase($1, $2, $3, $4)", id, p.Date, p.Amount, p.Description)
	if err != nil {
		return "", fmt.Errorf("wex.save_purchase execution failed: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	r.zl.Info("repo.SavePurchase: done")
	return id, nil
}

func (r *dbPurchaseRepository) GetPurchase(ctx context.Context, id string) (*model.Transaction, error) {
	r.zl.Info("repo.GetPurchase: started")

	// start db tx
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(ctx, "SELECT wex.get_purchase($1)", id)
	if err != nil {
		return nil, fmt.Errorf("wex.get_purchase execution failed: %v", err)
	}
	defer rows.Close()

	ts := model.Transaction{}
	err = rows.Scan(&ts.Date, &ts.Amount, &ts.Description)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	r.zl.Info("repo.GetPurchase: done")
	return &ts, nil
}

func (r *dbPurchaseRepository) Close() error {
	if r.pool != nil {
		r.pool.Close()
	}
	return nil
}
