package pgx

import (
	"context"
	"io"

	"github.com/jasonkwh/wex-test/internal/data/model"
)

type PurchaseRepository interface {
	SavePurchase(ctx context.Context, purchase model.Transaction) (id string, err error)
	GetPurchase(ctx context.Context, id string) (purchase model.Transaction, err error)

	io.Closer
}
