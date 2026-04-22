package main

import "fmt"

type Order struct {
	ID     string
	Amount float64
}

// ProcessOrder only knows about the PaymentMethod interface
// It has NO idea if it's CreditCard, PayPal, or Crypto
func ProcessOrder(order Order, payment PaymentMethod) {
	fmt.Printf("\n=== Processing Order %s ===\n", order.ID)
	fmt.Printf("Using: %s\n", payment.Name())

	if err := payment.Charge(order.Amount); err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		return
	}
	fmt.Println("✅ Order complete!")
}

// Process MULTIPLE payments — slice of interfaces!
func ProcessAll(orders []Order, methods []PaymentMethod) {
	for i, order := range orders {
		ProcessOrder(order, methods[i%len(methods)])
	}
}
