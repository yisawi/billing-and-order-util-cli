package billing

import "fmt"

func evalTier(remainingBalance float64) (string, float64) {
	tier := "Regular Customer"
	shippingDiscount := 0.0

	switch {
	case remainingBalance >= 500.0:
		tier = "Premium Customer"
		shippingDiscount = 15.0
	case remainingBalance >= 200.0:
		tier = "Gold Customer"
		shippingDiscount = 5.0
	}
	return tier, shippingDiscount
}

func ShowReceiptWithBalance(basePrice, userBalance float64) float64 {
	defer fmt.Println("Thank you for using our Billing CLI!.")
	defer fmt.Println()

	vatAmount, serviceAmount, totalPrice := calculateCharges(basePrice)

	if userBalance < totalPrice {
		fmt.Printf("❌ Transaction Declined: Insufficient balance!\n")
		fmt.Printf("Required: $%.1f | Your Balance: $%.1f\n", totalPrice, userBalance)
		fmt.Println("--------------------------------")
		return 0.0
	}

	remainingBalance := userBalance - totalPrice
	tier, shippingDiscount := evalTier(remainingBalance)

	fmt.Println("-----------------------")
	fmt.Println("--- RECEIPT DETAILS ---")
	fmt.Printf("Base Price:   $%.1f\n", basePrice)
	fmt.Printf("VAT (15%%):    $%.1f\n", vatAmount)
	fmt.Printf("Service (5%%): $%.1f\n", serviceAmount)
	fmt.Println("-----------------------")
	fmt.Printf("Total Bill:   $%.1f\n", totalPrice)
	fmt.Println("--- WALLET DETAILS ---")
	fmt.Printf("Starting Balance:   $%.1f\n", userBalance)
	fmt.Printf("Customer Tier:      %s\n", tier)
	if shippingDiscount > 0 {
		fmt.Printf("Congrats: You saved $%.1f on shipping!\n", shippingDiscount)
	}
	fmt.Printf("Remaining Balance:  $%.1f\n", (remainingBalance + shippingDiscount))
	fmt.Println("-----------------------")

	return totalPrice
}
