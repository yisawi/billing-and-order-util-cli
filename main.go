package main

import (
	"billing-and-order-util-cli/billing"
	"fmt"
)

func main() {
	fmt.Println("Starting Billing & Order Utility CLI...")
	fmt.Println()

	trackSales := billing.CreateSalesTracker()

	// --------------------------------------------------
	fmt.Println("[TEST 1: Successful Purchase]")
	price1 := 100.0
	actualPaid1 := billing.ShowReceiptWithBalance(price1, 650.0)

	if actualPaid1 > 0 {
		orders, total := trackSales(actualPaid1)
		fmt.Printf("Live Store Stats -> Orders Count: %d | Total Store Revenue: $%.1f\n\n", orders, total)
	}

	// --------------------------------------------------
	fmt.Println("[TEST 2: Failed Purchase due to Insufficient Funds]")
	price2 := 300.0
	actualPaid2 := billing.ShowReceiptWithBalance(price2, 100.0)

	if actualPaid2 > 0 {
		orders, total := trackSales(actualPaid2)
		fmt.Printf("Live Store Stats -> Orders: %d | Revenue: $%.1f\n\n", orders, total)
	} else {
		fmt.Println("⚠️ Stats Not Updated (Failed order is not tracked).\n")
	}

	// --------------------------------------------------
	fmt.Println("[TEST 3: Another Successful Purchase]")
	price3 := 80.0
	actualPaid3 := billing.ShowReceiptWithBalance(price3, 200.0)

	if actualPaid3 > 0 {
		orders, total := trackSales(actualPaid3)
		fmt.Printf("Live Store Stats -> Orders Count: %d | Total Store Revenue: $%.1f\n\n", orders, total)
	}
}
