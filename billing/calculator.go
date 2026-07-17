/*
Used for calcuate Tax and I implemented it with "Currying" Concept
*/
package billing

func TaxCalculator(rate float64) func(float64) float64 {
	return func(basePrice float64) float64 {
		return basePrice * rate
	}
}

var CalcVAT = TaxCalculator(0.15)
var CalcService = TaxCalculator(0.05)

func calculateCharges(basePrice float64) (vat, service, total float64) {
	vat = CalcVAT(basePrice)
	service = CalcService(basePrice)
	total = basePrice + vat + service
	return // naked return

}
