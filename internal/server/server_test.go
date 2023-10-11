package server

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
	"github.com/jasonkwh/wex-test/internal/data/model"
	"github.com/jasonkwh/wex-test/internal/exchangerate"
	"github.com/jasonkwh/wex-test/test/mocks"
	"github.com/onsi/gomega"
	"go.uber.org/zap"
)

func TestServer_SavePurchaseTransaction(t *testing.T) {
	type args struct {
		req           *purchasev1.SavePurchaseRequest
		transaction   *model.Transaction
		transactionId string
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
				transactionId: "test-transaction-id",
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

			tm.mockPurRepo.EXPECT().SavePurchase(gomock.Any(), tt.args.transaction).Return(tt.args.transactionId, nil)

			resp, err := srv.SavePurchaseTransaction(context.Background(), tt.args.req)
			g.Expect(err).To(gomega.BeNil())
			g.Expect(resp).To(gomega.Equal(tt.want))
		})
	}
}

func Test_server_GetPurchaseTransaction(t *testing.T) {
	type args struct {
		req         *purchasev1.GetPurchaseRequest
		transaction *model.Transaction
		httpResp    *http.Response
	}
	tests := []struct {
		name string
		args args
		want *purchasev1.GetPurchaseResponse
	}{
		{
			name: "test normal get purchase",
			args: args{
				req: &purchasev1.GetPurchaseRequest{
					Id:       "test-transaction-id",
					Currency: "Dollar",
				},
				transaction: &model.Transaction{
					Description: "test transaction",
					Date:        "2023-06-01",
					Amount:      1234,
				},
				httpResp: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"data": [{"exchange_rate": "1.494","record_date": "2023-06-01"}]}`))),
				},
			},
			want: &purchasev1.GetPurchaseResponse{
				Description: "test transaction",
				TransactionDate: &purchasev1.Date{
					Year:  2023,
					Month: 6,
					Day:   1,
				},
				Amount:          "12.34",
				Id:              "test-transaction-id",
				ExchangeRate:    "1.494",
				ConvertedAmount: "18.44",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)
			srv, tm := newMockServer(ctrl)

			tm.mockPurRepo.EXPECT().GetPurchase(gomock.Any(), tt.args.req.Id).Return(tt.args.transaction, nil)
			tm.mockHttpClient.EXPECT().Do(gomock.Any()).Return(tt.args.httpResp, nil)

			resp, err := srv.GetPurchaseTransaction(context.Background(), tt.args.req)
			g.Expect(err).To(gomega.BeNil())
			g.Expect(resp).To(gomega.Equal(tt.want))
		})
	}
}

type testMocks struct {
	mockPurRepo    *mocks.MockPurchaseRepository
	mockHttpClient *mocks.MockHTTPClient
}

func newMockServer(ctrl *gomock.Controller) (*server, *testMocks) {
	tm := testMocks{
		mockPurRepo:    mocks.NewMockPurchaseRepository(ctrl),
		mockHttpClient: mocks.NewMockHTTPClient(ctrl),
	}
	zl, _ := zap.NewDevelopment()

	srv := server{
		repo: tm.mockPurRepo,
		ret:  exchangerate.NewRetriever(tm.mockHttpClient, 6),
		zl:   zl,
	}

	return &srv, &tm
}
