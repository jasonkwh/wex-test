package server

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
	"github.com/jasonkwh/wex-test/internal/data/model"
	"github.com/jasonkwh/wex-test/test/mocks"
	"github.com/onsi/gomega"
	"go.uber.org/zap"
)

func TestServer_SavePurchaseTransaction(t *testing.T) {
	type args struct {
		req         *purchasev1.SavePurchaseRequest
		transaction *model.Transaction
		id          string
	}
	tests := []struct {
		name string
		args args
		want *purchasev1.SavePurchaseResponse
	}{
		{
			name: "test normal save purchase transaction",
			args: args{
				req: &purchasev1.SavePurchaseRequest{
					Description: "test description",
					TransactionDate: &purchasev1.Date{
						Year:  2023,
						Month: 6,
						Day:   1,
					},
					Amount: "12.34",
				},
				transaction: &model.Transaction{
					Description: "test description",
					Date:        "2023-06-01",
					Amount:      1234,
				},
				id: "test-transaction-id",
			},
			want: &purchasev1.SavePurchaseResponse{
				Id: "test-transaction-id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)
			srv, tm := newMockServer(ctrl)

			tm.mockPurRepo.EXPECT().SavePurchase(gomock.Any(), tt.args.transaction).Return(tt.args.id, nil)

			resp, err := srv.SavePurchaseTransaction(context.Background(), tt.args.req)
			g.Expect(err).To(gomega.BeNil())
			g.Expect(resp).To(gomega.Equal(tt.want))
		})
	}
}

type testMocks struct {
	mockPurRepo *mocks.MockPurchaseRepository
}

func newMockServer(ctrl *gomock.Controller) (*server, *testMocks) {
	tm := testMocks{
		mockPurRepo: mocks.NewMockPurchaseRepository(ctrl),
	}
	zl, _ := zap.NewDevelopment()

	srv := server{
		within: 6,
		repo:   tm.mockPurRepo,
		zl:     zl,
	}

	return &srv, &tm
}
