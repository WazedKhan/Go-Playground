package main

import "fmt"

// func ProcessCreditCard(c CreditCard) { c.Charge(100) }
// func ProcessPayPal(p PayPal)         { p.Charge(100) }
// func ProcessCrypto(c Crypto)         { c.Charge(100) }
// Every new payment method = new function 😭

// ✅ WITH interfaces — one function rules them all
type PaymentMethod interface {
    Charge(amount float64) error
}

func ProcessPayment(p PaymentMethod, amount float64) {
    if err := p.Charge(amount); err != nil {
        fmt.Println("Payment failed:", err)
        return
    }
    fmt.Println("Payment successful!")
}
