package main

// PaymentMethod is the contract every payment type must fulfill
type PaymentMethod interface {
	Charge(amount float64) error
	Refund(amount float64) error
	Name() string
}
