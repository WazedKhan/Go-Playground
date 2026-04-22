package main

import "fmt"

type CreditCard struct {
	CardNumber string
	Balance    float64
}

func (c *CreditCard) Charge(amount float64) error {
	if c.Balance < amount {
		return fmt.Errorf("insufficient balance: have %.2f, need %.2f", c.Balance, amount)
	}
	c.Balance -= amount
	fmt.Printf("[CreditCard] Charged $%.2f. Remaining: $%.2f\n", amount, c.Balance)
	return nil
}

func (c *CreditCard) Refund(amount float64) error {
	c.Balance += amount
	fmt.Printf("[CreditCard] Refunded $%.2f\n", amount)
	return nil
}

func (c *CreditCard) Name() string {
	return fmt.Sprintf("CreditCard ****%s", c.CardNumber[len(c.CardNumber)-4:])
}
