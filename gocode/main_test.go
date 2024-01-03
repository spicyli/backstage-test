package main

import (
	"testing"
)
func TestFunc(t *testing.T) {
	t.Run("Valid Min", func(t *testing.T) {
		somelist := [6]int{2, 3, 5, 7, 11, 1}
		result := MinOfInts(somelist)
		expected := 1
		if result != expected {
			t.Errorf("Incorrect Min Value.\nExpected: %v.\nActual: %v\n", expected, result)
		}
	})

	t.Run("Valid Max", func(t *testing.T) {
		somelist := [6]int{2, 3, 5, 7, 11, 1}
		result := MaxOfInts(somelist)
		expected := 11
		if result != expected {
			t.Errorf("Incorrect Max Value.\nExpected: %v.\nActual: %v\n", expected, result)
		}
	})

	t.Run("Valid uppercase", func(t *testing.T) {
		somestring := "somestring"
		result := StringToLower(somestring)
		expected := "SOMESTRING"
		if result != expected {
			t.Errorf("Incorrect string Value.\nExpected: %v.\nActual: %v\n", expected, result)
		}
	})
}
