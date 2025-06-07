package factory_method

import (
	"errors"
	"fmt"
)

const (
	Cash = iota + 1
	DebitCard
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(t int) (PaymentMethod, error) {
	switch t {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", t))
	}
}

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

type DebitCardPM struct{}

func (d *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

type CreditCardPM struct{}

func (c *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using new credit card implementation\n", amount)
}
