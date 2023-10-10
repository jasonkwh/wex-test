//go:build integration
// +build integration

package exchangerate_test

import (
	"testing"

	"github.com/jasonkwh/wex-test/internal/data/model"
	"github.com/jasonkwh/wex-test/internal/exchangerate"
	"github.com/onsi/gomega"
)

func TestRetrieve(t *testing.T) {
	type args struct {
		date     string
		currency string
		within   int
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
			},
			want: &model.ExchangeRate{
				ExchangeRate: "6.897",
				RecordDate:   "2022-12-31",
			},
			wantErr: nil,
		},
		{
			name: "test normal retrieve (currency = Dollar)",
			args: args{
				date:     "2023-03-31",
				currency: "Dollar",
				within:   6,
			},
			want: &model.ExchangeRate{
				ExchangeRate: "1.494",
				RecordDate:   "2023-03-31",
			},
			wantErr: nil,
		},
		{
			name: "test if exchange rate is NOT valid (within 1 month)",
			args: args{
				date:     "2023-03-01",
				currency: "Renminbi",
				within:   1,
			},
			want:    nil,
			wantErr: exchangerate.ErrFailedRetrieve,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)

			// execute retrieve
			exr, err := exchangerate.Retrieve(tt.args.date, tt.args.currency, tt.args.within)

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
