package billing

import "fmt"

const VatRate = 0.15 // just a mock data
const ServiceTaxRate = 0.15 // just a mock data

// take the base price and print the recepit
func ShowInitialReceipt(basePrice float64) {
	vatAmount := basePrice * VatRate
	serviceAmount := basePrice * ServiceTaxRate
	totalPrice := basePrice + vatAmount + serviceAmount


	fmt.Println("--- RECEIPT DETAILS ---")
	fmt.Printf("Base Price:   $%.1f\n", basePrice)
	fmt.Printf("VAT (15%%):    $%.1f\n", vatAmount)
	fmt.Printf("Service (15%%): $%.1f\n", serviceAmount)
	fmt.Println("-----------------------")
	fmt.Printf("Total Bill:   $%.1f\n", totalPrice)
	fmt.Println("-----------------------")

}