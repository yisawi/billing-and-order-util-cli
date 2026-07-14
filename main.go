package main

import (

	"billing-and-order-util-cli/billing"
	"fmt"
)

func main() {
	fmt.Println("Starting Billing & Order Utility CLI...")
	fmt.Println()


	samplePrice := 169.169

	billing.ShowInitialReceipt(samplePrice)
}