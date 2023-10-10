package utils

import (
	"fmt"
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
