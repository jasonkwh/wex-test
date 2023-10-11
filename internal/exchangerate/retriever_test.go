package exchangerate_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jasonkwh/wex-test/internal/data/model"
	"github.com/jasonkwh/wex-test/internal/exchangerate"
	"github.com/jasonkwh/wex-test/test/mocks"
	"github.com/onsi/gomega"
)

func TestRetriever_Get(t *testing.T) {
	type args struct {
		date     string
		currency string
		within   int
		httpResp *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *model.ExchangeRate
		wantErr error
	}{
		{
			name: "test normal retrieve (currency = Renminbi)",
			args: args{
				date:     "2022-12-31",
				currency: "Renminbi",
				within:   6,
				httpResp: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"data": [{"exchange_rate": "6.897","record_date": "2022-12-31"}]}`))),
				},
			},
			want: &model.ExchangeRate{
				ExchangeRate: "6.897",
				RecordDate:   "2022-12-31",
			},
			wantErr: nil,
		},
		{
			name: "test if exchange rate is NOT valid (within 1 month)",
			args: args{
				date:     "2022-03-01",
				currency: "Renminbi",
				within:   1,
				httpResp: &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(`{"data": [{"exchange_rate": "6.373","record_date": "2021-12-31"}]}`))),
				},
			},
			want:    nil,
			wantErr: exchangerate.ErrFailedRetrieve,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)
			ret, tm := newMockRetriever(ctrl, tt.args.within)

			tm.mockHttpClient.EXPECT().Do(gomock.Any()).Return(tt.args.httpResp, nil)

			exr, err := ret.Get(tt.args.date, tt.args.currency)

			// expecting error
			if tt.wantErr == nil {
				g.Expect(err).To(gomega.BeNil())
			} else {
				g.Expect(err).To(gomega.Equal(tt.wantErr))
			}

			// expecting result
			if tt.want == nil {
				g.Expect(exr).To(gomega.BeNil())
			} else {
				g.Expect(exr).To(gomega.Equal(tt.want))
			}
		})
	}
}

type testMocks struct {
	mockHttpClient *mocks.MockHTTPClient
}

func newMockRetriever(ctrl *gomock.Controller, within int) (*exchangerate.Retriever, *testMocks) {
	tm := testMocks{
		mockHttpClient: mocks.NewMockHTTPClient(ctrl),
	}

	return exchangerate.NewRetriever(tm.mockHttpClient, within), &tm
}
