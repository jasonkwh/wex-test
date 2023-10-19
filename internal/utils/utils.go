package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
)

func ToFormattedDate(date *purchasev1.Date) string {
	// to dealt with months or days less than 10
	month := date.Month
	if len(month) == 1 {
		month = fmt.Sprintf("0%s", month)
	}
	day := date.Day
	if len(day) == 1 {
		day = fmt.Sprintf("0%s", day)
	}

	return fmt.Sprintf("%s-%s-%s", date.Year, month, day)
}

func ToUpstreamDate(date string) *purchasev1.Date {
	dateParts := strings.Split(date, "-")

	return &purchasev1.Date{
		Year:  dateParts[0],
		Month: dateParts[1],
		Day:   dateParts[2],
	}
}

// exchange rate conversion
func GetConvertedAmount(amount float64, exchangeRate string) float64 {
	rate, err := strconv.ParseFloat(exchangeRate, 64)
	if err != nil {
		return 0
	}
	return amount * rate
}
