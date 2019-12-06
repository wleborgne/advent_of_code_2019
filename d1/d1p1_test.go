package main

import (
	"testing"
)

func TestCalculateUnit(t *testing.T) {
	expected := 654
	actual := calculate_unit("1969")

	if actual != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, actual)
	}
}
