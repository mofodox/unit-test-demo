package listing

// Listing object
type Listing struct {
	Id    string
	Name  string
	Price float64
}

func CalculateCommission(price float64) (commission float64) {
	// Returns the commission
	return price * 0.03
}
