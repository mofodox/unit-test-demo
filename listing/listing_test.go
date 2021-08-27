package listing

import "testing"

func TestCalculateCommission(t *testing.T) {
	// Arrange your objects
	listing := Listing{
		Id: "1",
		Name: "11 Jackson Square",
		Price: 1000000.00,
	}

	// Act on the object
	commission := CalculateCommission(listing.Price)

	// comission = price * 0.03

	expectedValue := 30000.00

	// Assert
	if commission != expectedValue {
		t.Errorf("commission %.2f, expected commission %.2f", commission, expectedValue)
	}
}

type TestCalculateMultipleListingCommission struct {
	listing Listing
	expectedCommission float64
}

func TestTDT(t *testing.T) {
	listingCases := []TestCalculateMultipleListingCommission{
		{
			listing: Listing{
				Id:    "1",
				Name:  "11 Jackson Square",
				Price: 1000000.00,
			},
			expectedCommission: 30000.00,
		},
		{
			listing: Listing{
				Id:    "2",
				Name:  "255 Toa Payoh",
				Price: 2000000.00,
			},
			expectedCommission: 60000.00,
		},
		{
			listing: Listing{
				Id:    "3",
				Name:  "12 Bukit Merah Road",
				Price: 3000000.00,
			},
			expectedCommission: 90000.00,
		},
	}

	for _, listing := range listingCases {
		if commission := CalculateCommission(listing.listing.Price); commission != listing.expectedCommission {
			t.Errorf("commission %.2f expected commission %.2f listing price %.2f", commission, listing.expectedCommission, listing.listing.Price)
		}
	}
}
