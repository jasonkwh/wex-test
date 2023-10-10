package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
)

func ToFormattedDate(date *purchasev1.Date) string {
	// to dealt with months or days less than 10
	month := fmt.Sprintf("%d", date.Month)
	if date.Month < 10 {
		month = fmt.Sprintf("0%d", date.Month)
	}
	day := fmt.Sprintf("%d", date.Day)
	if date.Day < 10 {
		month = fmt.Sprintf("0%d", date.Day)
	}

	return fmt.Sprintf("%d-%s-%s", date.Year, month, day)
}

func ToUpstreamDate(date string) *purchasev1.Date {
	dateParts := strings.Split(date, "-")
	year, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	day, _ := strconv.Atoi(dateParts[2])

	return &purchasev1.Date{
		Year:  int32(year),
		Month: int32(month),
		Day:   int32(day),
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
