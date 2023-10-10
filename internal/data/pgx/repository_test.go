//go:build integration
// +build integration

package pgx_test

import (
	"context"
	"testing"

	"github.com/jasonkwh/wex-test/internal/config"
	"github.com/jasonkwh/wex-test/internal/data/model"
	"github.com/jasonkwh/wex-test/internal/data/pgx"
	"github.com/onsi/gomega"
	"go.uber.org/zap"
)

func TestRepository(t *testing.T) {
	tests := []struct {
		name       string
		arg        *model.Transaction
		wantAmount int
	}{
		{
			name: "test normal save & get",
			arg: &model.Transaction{
				Description: "test transaction",
				Date:        "2023-05-01",
				Amount:      2560, // $25.60
			},
			wantAmount: 2560,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctx := context.Background()
			zl, err := zap.NewDevelopment()
			if err != nil {
				t.Fatal(err)
			}

			rp, err := connectDb(zl)
			if err != nil {
				t.Fatal(err)
			}

			id, err := rp.SavePurchase(ctx, tt.arg)
			g.Expect(err).To(gomega.BeNil())
			g.Expect(id).To(gomega.Not(gomega.BeEmpty()))

			resp, err := rp.GetPurchase(ctx, id)
			g.Expect(err).To(gomega.BeNil())
			g.Expect(resp.Amount).To(gomega.Equal(tt.wantAmount))
		})
	}
}

func connectDb(zl *zap.Logger) (pgx.PurchaseRepository, error) {
	cfg := config.DatabaseConfig{
		User:     "wex_dev",
		Password: "password",
		Host:     "127.0.0.1",
		Database: "postgres",
		Port:     5432,
	}

	return pgx.CreatePurchaseRepository(cfg, zl)
}
