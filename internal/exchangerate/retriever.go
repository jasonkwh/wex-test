package exchangerate

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/jasonkwh/wex-test/internal/data/model"
	"go.uber.org/multierr"
)

var ErrFailedRetrieve = fmt.Errorf("purchase cannot converted to the target currency")

func Retrieve(date, currency string, within int) (*model.ExchangeRate, error) {
	resp, err := http.Get(fmt.Sprintf(
		"https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?filter=record_date:lte:%s,currency:eq:%s&format=json&sort=-record_date&fields=exchange_rate,record_date&page[number]=1&page[size]=1",
		date,
		currency,
	))
	if err != nil {
		err = multierr.Append(ErrFailedRetrieve, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = multierr.Append(ErrFailedRetrieve, err)
		return nil, err
	}

	exrs := model.ExchangeRates{}
	err = json.Unmarshal(body, &exrs)
	if err != nil {
		err = multierr.Append(ErrFailedRetrieve, err)
		return nil, err
	}

	// we only retrieve one record
	exr := exrs.Data[0]

	if !isValid(date, exr.RecordDate, within) {
		return nil, ErrFailedRetrieve
	}
	return &exr, nil
}

func isValid(date1, date2 string, within int) bool {
	diff, err := getMonthsDifference(date1, date2)
	if err != nil {
		return false
	}
	if diff > within {
		return false
	}
	return true
}

func getMonthsDifference(date1, date2 string) (int, error) {
	// parse
	layout := "2006-01-02"
	d1, err := time.Parse(layout, date1)
	if err != nil {
		return 0, err
	}
	d2, err := time.Parse(layout, date2)
	if err != nil {
		return 0, err
	}

	// get diffs
	yearDiff := d1.Year() - d2.Year()
	monthDiff := int(d1.Month()) - int(d2.Month())

	// for negative month difference
	if monthDiff < 0 {
		monthDiff += 12
		yearDiff--
	}

	return int(math.Abs(float64(yearDiff*12 + monthDiff))), nil
}
