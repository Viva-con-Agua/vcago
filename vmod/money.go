package vmod

// Money represent the simple Money object. The Amount is a cent representation.
type Money struct {
	Amount   int64  `bson:"amount" json:"amount"`
	Currency string `bson:"currency" json:"currency" validate:"required"`
}

/*
func (i *Money) ValidateAmount(minAmount int64) (err error) {
	vErr := new(ValidationError)
	if i.Amount < minAmount {
		vErr.Errors = []string{"Amount is to low"}
		return vErr
	}
	return
}

func (i *Money) ValidateCurrency(currency string) (err error) {
	if !strings.Contains(currency, i.Currency) {
		vErr := new(ValidationError)
		vErr.Errors = []string{"Currency is not supported!"}
		return vErr
	}
	return nil
}*/
