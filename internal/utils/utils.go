package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
)

func ToFormattedDate(date *purchasev1.Date) string {
	return fmt.Sprintf("%d-%d-%d", date.Year, date.Month, date.Day)
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
	return math.Round(amount*rate*100) / 100
}
