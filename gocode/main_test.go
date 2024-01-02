package main

import (
	"testing"
)
func TestMinFunc(t *testing.T) {
	t.Run("Valid Min", func(t *testing.T) {
		somelist := [6]int{2, 3, 5, 7, 11, 1}
		result := MinOfInts(somelist)
		expected := 1
		if result != expected {
			t.Errorf("Incorrect Min Value.\nExpected: %v.\nActual: %v\n", expected, result)
		}
	})
}
