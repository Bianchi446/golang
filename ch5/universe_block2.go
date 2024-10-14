package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	// Return the division result, remainder, and nil as no error occurred
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	// Print the result with correct number of return values
	result, remainder, err := divAndRemainder(25, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
	}
}
