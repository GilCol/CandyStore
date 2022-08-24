package main

import (
	"reflect"
	"testing"
)

var c1 = Customer{
	Name:           "Erik",
	FavouriteSnack: "Chips",
	TotalSnacks:    2,
}

var c2 = Customer{
	Name:           "Johan",
	FavouriteSnack: "Nötter",
	TotalSnacks:    4,
}

var c3 = Customer{
	Name:           "Daniel",
	FavouriteSnack: "Banan",
	TotalSnacks:    6,
}

var c4 = Customer{
	Name:           "Hampus",
	FavouriteSnack: "Godis",
	TotalSnacks:    8,
}

func Test_CustomerSort_Given_Valid_Data(t *testing.T) {
	var topCustomers []Customer
	topCustomers = append(topCustomers, c3)
	topCustomers = append(topCustomers, c2)
	topCustomers = append(topCustomers, c1)
	topCustomers = append(topCustomers, c4)

	sortCustomers(topCustomers)

	var expectedSortedCustomers []Customer
	expectedSortedCustomers = append(expectedSortedCustomers, c4)
	expectedSortedCustomers = append(expectedSortedCustomers, c3)
	expectedSortedCustomers = append(expectedSortedCustomers, c2)
	expectedSortedCustomers = append(expectedSortedCustomers, c1)

	got := topCustomers
	want := expectedSortedCustomers

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test_CustomerSort_Given_Empty_Data(t *testing.T) {
	var topCustomers []Customer

	sortCustomers(topCustomers)

	var expectedSortedCustomers []Customer

	got := topCustomers
	want := expectedSortedCustomers

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
