package main 

import (
	"fmt"
	"errors"
)


func divAndRemainder(numerator int, denominator int) (result int, remainder int, err error) {
	result, remainder = 20, 10
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return result, remainder, err
	}
	return numerator / denominator, numerator % denominator, nil

}


func main() {
	result, remainder, err := divAndRemainder(20, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
	}
}