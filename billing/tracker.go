/*
Sales Tracker : Responsible only for memory management and tracking sales figures.

also implemented with "Clousers" Concept
*/
package billing

func CreateSalesTracker() func(float64) (int, float64) {
	totalSales := 0.0
	successfulOrders := 0

	return func(orderAmount float64) (int, float64) {
		successfulOrders++
		totalSales += orderAmount
		return successfulOrders, totalSales
	}
}