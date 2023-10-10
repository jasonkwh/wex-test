package server

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
	"github.com/jasonkwh/wex-test/internal/utils"
	"go.uber.org/multierr"
)

func ValidateSavePurchase(req any) error {
	r, ok := req.(*purchasev1.SavePurchaseRequest)
	if !ok {
		return ErrInvalidType
	}
	return validateTransactionDetails(r)
}

func validateTransactionDetails(r *purchasev1.SavePurchaseRequest) error {
	errs := validation.ValidateStruct(r,
		validation.Field(r.Description, validation.Length(0, 50)),
		validation.Field(r.Amount, validation.Min(1)), //minimum amount is $0.01
	)
	derr := validation.Validate(
		utils.ToFormattedDate(r.TransactionDate),
		validation.Date("2006-01-02"),
	)

	errs = multierr.Append(errs, derr)
	return errs
}
