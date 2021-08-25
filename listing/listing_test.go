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
		t.Errorf("commission %.2f expected commission %.2f listing price %.2f", commission, expectedCommission, listing.Price)
	}
}

func TestCalculatePriceAfterCommission(t *testing.T) {
	listing := Listing{
		Id:    "1",
		Name:  "11 Jackson Square",
		Price: 1000000.00,
	}

	priceAfterCommission := CalculatePriceAfterCommission(listing.Price)
	expectedPriceAfterCommission := 970000.00

	if priceAfterCommission != expectedPriceAfterCommission {
		t.Errorf("price after commission %.2f expected price after commission %.2f listing price %.2f", priceAfterCommission, expectedPriceAfterCommission, listing.Price)
	}
}

/********* Table Driven Tests *********/

type TestMultipleListingPrice struct {
	listing                      Listing
	expectedPriceAfterCommission float64
}

func TestTDT(t *testing.T) {
	listingCases := []TestMultipleListingPrice{
		{
			listing: Listing{
				Id:    "1",
				Name:  "11 Jackson Square",
				Price: 1000000.00,
			},
			expectedPriceAfterCommission: 970000.00,
		},
		{
			listing: Listing{
				Id:    "2",
				Name:  "255 Toa Payoh",
				Price: 550000.00,
			},
			expectedPriceAfterCommission: 533500.00,
		},
		{
			listing: Listing{
				Id:    "3",
				Name:  "12 Bukit Merah Road",
				Price: 650400.00,
			},
			expectedPriceAfterCommission: 630888.00,
		},
	}

	for _, listing := range listingCases {
		if priceAfterCommission := CalculatePriceAfterCommission(listing.listing.Price); priceAfterCommission != listing.expectedPriceAfterCommission {
			t.Errorf("listing price after commission %.2f expectedPriceAfterCommission %.2f listing price %.2f", priceAfterCommission, listing.expectedPriceAfterCommission, listing.listing.Price)
		}
	}
}
