package main

import (
	"flag"
	"fmt"
	"payment-processor/cmd"
	"payment-processor/provider"
)

const (
	BKASH  = "bkash"
	NAGAD  = "nagad"
	ROCKET = "rocket"
)

func printRes(res cmd.PaymentResult, providerTag string) {
	percentage := (res.Fee / res.Amount) * 100
	fmt.Printf("💰 %s Payment Initiated \n", providerTag)
	fmt.Printf("   Amount  : %.2f BDT\n", res.Amount)
	fmt.Printf("   Fee     : %.2f BDT (%.2f%%)\n", res.Fee, percentage)
	fmt.Printf("   Total   : %.2f BDT\n", res.Total)
	fmt.Println("   Status  : Success ✓")
}

func main() {
	amount := flag.Float64("amount", 0, "")
	providerName := flag.String("provider", "", "")
	flag.Parse()


	var p cmd.PaymentProvider
	switch *providerName {
	case BKASH:
		p = provider.Bkash{}

	case NAGAD:
		p = provider.Nagad{}

	case ROCKET:
		p = provider.Rocket{}
	default:
		fmt.Println("unknown provider")
    	return
	}

	res, err := p.Pay(*amount)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	printRes(res, *providerName)
}
