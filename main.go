package main

import (

	"billing-and-order-util-cli/billing"
	"fmt"
)

func main() {
	fmt.Println("Starting Billing & Order Utility CLI...")
	fmt.Println()


	price := 100.0

	fmt.Println("[TEST 1: Successful Purchase]")
	billing.ShowReceiptWithBalance(price, 650.0)
	fmt.Println()
}