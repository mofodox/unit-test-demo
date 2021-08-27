package listing

type Listing struct {
	Id string
	Name string
	Price float64
}

func CalculateCommission(price float64) float64 {
	return price * 0.03
}
