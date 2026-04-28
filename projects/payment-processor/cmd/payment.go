package cmd

type PaymentProvider interface {
	Pay(amount float64) (PaymentResult, error)
}

type PaymentResult struct {
	Amount float64
	Fee    float64
	Total  float64
	Status string
}
