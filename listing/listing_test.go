package listing_test

import (
	"testing"

	"unit-test-demo/listing"
)

func TestCalculateCommission(t *testing.T) {
	listingStruct := listing.Listing{
		Id:    "1",
		Name:  "11 Jackson Square",
		Price: 1000000.00,
	}

	// Find the commission by using the CalculateCommission function
	commission := listing.CalculateCommission(listingStruct.Price)
	// We know that to find the commission of a listing, the formula is: price * 0.03
	expectedCommission := 30000.00

	// Assert
	if commission != expectedCommission {
		t.Errorf("commission %.2f expected commission %.2f listing price %.2f", commission, expectedCommission, listingStruct.Price)
	}
}

func TestCalculatePriceAfterCommission(t *testing.T) {
	listingStruct := listing.Listing{
		Id:    "1",
		Name:  "11 Jackson Square",
		Price: 1000000.00,
	}

	priceAfterCommission := listing.CalculatePriceAfterCommission(listingStruct.Price)
	expectedPriceAfterCommission := 970000.00

	if priceAfterCommission != expectedPriceAfterCommission {
		t.Errorf("price after commission %.2f expected price after commission %.2f listing price %.2f", priceAfterCommission, expectedPriceAfterCommission, listingStruct.Price)
	}
}

/********* Table Driven Tests *********/

type TestMultipleListingPrice struct {
	listingStruct                listing.Listing
	expectedPriceAfterCommission float64
}

func TestTDT(t *testing.T) {
	listingCases := []TestMultipleListingPrice{
		{
			listingStruct: listing.Listing{
				Id:    "1",
				Name:  "11 Jackson Square",
				Price: 1000000.00,
			},
			expectedPriceAfterCommission: 970000.00,
		},
		{
			listingStruct: listing.Listing{
				Id:    "2",
				Name:  "255 Toa Payoh",
				Price: 550000.00,
			},
			expectedPriceAfterCommission: 533500.00,
		},
		{
			listingStruct: listing.Listing{
				Id:    "3",
				Name:  "12 Bukit Merah Road",
				Price: 650400.00,
			},
			expectedPriceAfterCommission: 630888.00,
		},
	}

	for _, l := range listingCases {
		if priceAfterCommission := listing.CalculatePriceAfterCommission(l.listingStruct.Price); priceAfterCommission != l.expectedPriceAfterCommission {
			t.Errorf("listing price after commission %.2f expectedPriceAfterCommission %.2f listing price %.2f", priceAfterCommission, l.expectedPriceAfterCommission, l.listingStruct.Price)
		}
	}
}
