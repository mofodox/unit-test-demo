package listing

import "testing"

func TestCalculateCommission(t *testing.T) {
	listing := Listing{
		Id:    "1",
		Name:  "11 Jackson Square",
		Price: 1000000.00,
	}

	// Find the commission by using the CalculateCommission function
	commission := CalculateCommission(listing.Price)
	// We know that to find the commission of a listing, the formula is: price * 0.03
	expectedCommission := 30000.00

	// Assert
	if commission != expectedCommission {
		t.Errorf("commission %.2f expected commission %.2f", commission, expectedCommission)
	}
}
