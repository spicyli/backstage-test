package main

import (
    "fmt"
	"strings"
)

func main() {
	somelist := [6]int{2, 3, 5, 7, 11, 1}
    fmt.Println(MinOfInts(somelist))
}

func MinOfInts(vars [6]int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOfInts(vars [6]int) int {
	max := vars[0]

	for _, i := range vars {
		if i > max {
			max = i
		}
	}

	return max
}

func StringToLower(word string) string {
	return strings.ToLower(word)
}