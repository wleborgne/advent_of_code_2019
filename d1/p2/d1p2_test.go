package main

import (
	"testing"
)

func TestCalculateCargo(t *testing.T) {
	expected := 966
	actual := calculate_cargo("1969")

	if actual != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, actual)
	}
}
