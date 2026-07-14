package billing

import "fmt"

const VatRate = 0.15 // just a mock data
const ServiceTaxRate = 0.05 // just a mock data

// take the base price and print the recepit
func ShowReceiptWithBalance(basePrice float64, userBalance float64) {
	vatAmount := basePrice * VatRate
	serviceAmount := basePrice * ServiceTaxRate
	totalPrice := basePrice + vatAmount + serviceAmount


	// check if the user has enough money
	if userBalance < totalPrice {
		fmt.Printf("❌ Transaction Declined: Insufficient balance!\n")
		fmt.Printf("Required: $%.1f | Your Balance: $%.1f\n", totalPrice, userBalance)
		fmt.Println("--------------------------------")
		return
	}
	
	// check the user level for discount
	remainingBalance := userBalance - totalPrice
	tier := ""
	shippingDiscount := 0.0

	switch {
	case remainingBalance >= 500.0:
		tier = "Premium Customer"
		shippingDiscount = 15.0
	case remainingBalance >= 200.0:
		tier = "Gold Customer"
		shippingDiscount = 5.0
	default:
		tier = "Regular Customer"
		shippingDiscount = 0.0
	}
	

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


}