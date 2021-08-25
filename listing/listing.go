package listing

// Listing object
type Listing struct {
	Id    string
	Name  string
	Price float64
}

func CalculateCommission(price float64) float64 {
	// Returns the commission
	return price * 0.03
}

func CalculatePriceAfterCommission(price float64) float64 {
	commission := CalculateCommission(price)
	result := price - commission

	return result
}
