//go:build integration
// +build integration

package integration

import (
	"context"
	"testing"
	"time"

	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
	"github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestIntegration(t *testing.T) {
	ctx := context.Background()
	g := gomega.NewWithT(t)

	// expect value
	amount := "25.60"
	exchangeRate := "1.494"

	// dial grpc
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	cl := purchasev1.NewPurchaseServiceClient(conn)

	// try to save a transaction
	saveResp, err := cl.SavePurchaseTransaction(ctx, &purchasev1.SavePurchaseRequest{
		Description: "test description",
		TransactionDate: &purchasev1.Date{
			Year:  "2023",
			Month: "03",
			Day:   "31",
		},
		Amount: amount,
	})
	g.Expect(err).To(gomega.BeNil())

	time.Sleep(2 * time.Second)

	// try to get the transaction back
	resp, err := cl.GetPurchaseTransaction(ctx, &purchasev1.GetPurchaseRequest{
		Id:       saveResp.Id,
		Currency: "Dollar",
	})
	g.Expect(err).To(gomega.BeNil())
	g.Expect(resp.Amount).To(gomega.Equal(amount))
	g.Expect(resp.ExchangeRate).To(gomega.Equal(exchangeRate))
}
