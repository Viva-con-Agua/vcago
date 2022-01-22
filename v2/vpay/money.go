package vpay

import (
	"strings"

	"github.com/Viva-con-Agua/vcago/verr"
)

type Money struct {
	Amount   int64  `bson:"amount" json:"amount" validate:"required"`
	Currency string `bson:"currency" json:"currency" validate:"required"`
}

func (i *Money) ValidateAmount(minAmount int64) (err error) {
	vErr := new(verr.ValidationError)
	if i.Amount < minAmount {
		vErr.Errors = []string{"Amount is to low"}
		return vErr
	}
	return
}

func (i *Money) ValidateCurrency(currency string) (err error) {
	if !strings.Contains(currency, i.Currency) {
		vErr := new(verr.ValidationError)
		vErr.Errors = []string{"Currency is not supported!"}
		return vErr
	}
	return nil
}
