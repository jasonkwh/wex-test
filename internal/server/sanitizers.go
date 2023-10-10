package server

import (
	"fmt"
	"strings"

	"github.com/jasonkwh/wex-test-upstream/svc/purchasev1"
)

var ErrInvalidType = fmt.Errorf("the type provided to the validator/sanitizer was invalid")
var ErrEmptyTransId = fmt.Errorf("the identifier provided to the validator/sanitizer was empty")

func SanitizeSavePurchase(req any) error {
	r, ok := req.(*purchasev1.SavePurchaseRequest)
	if !ok {
		return ErrInvalidType
	}
	if r.Description != "" {
		r.Description = strings.TrimSpace(r.Description)
	}
	return nil
}

func SanitizeGetPurchase(req any) error {
	r, ok := req.(*purchasev1.GetPurchaseRequest)
	if !ok {
		return ErrInvalidType
	}
	if r.Id == "" {
		return ErrEmptyTransId
	}
	r.Id = strings.TrimSpace(r.Id)
	return nil
}
