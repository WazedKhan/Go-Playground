package main

import "fmt"

func main() {
	// Create concrete types
	card := &CreditCard{CardNumber: "1234567890123456", Balance: 500}
	// pp   := &PayPal{Email: "user@example.com", Balance: 200}

	// Store them as interfaces — this is polymorphism!
	var methods []PaymentMethod = []PaymentMethod{card}

	orders := []Order{
		{ID: "ORD-001", Amount: 99.99},
		{ID: "ORD-002", Amount: 49.99},
		{ID: "ORD-003", Amount: 999.00}, // this will fail
	}

	ProcessAll(orders, methods)

	// Type assertion — check the real type at runtime
	for _, m := range methods {
		if cc, ok := m.(*CreditCard); ok {
			fmt.Printf("\nCredit card balance: $%.2f\n", cc.Balance)
		}
	}
}
