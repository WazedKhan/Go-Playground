package provider

import (
	"errors"
	"payment-processor/cmd"
)

type Rocket struct{}

func (r Rocket) Pay(amount float64) (cmd.PaymentResult, error) {
	if amount <= 0 {
		return cmd.PaymentResult{}, errors.New("amount must be greater than 0")
	}
	feePercentage := 1.5 / 100.0
	fee := amount * feePercentage

	return cmd.PaymentResult{
		Amount: amount,
		Fee:    fee,
		Total:  amount + fee,
		Status: "Success ✓",
	}, nil
}
