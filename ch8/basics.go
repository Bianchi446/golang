package main

import (
	"errors"
	"fmt"
	"os"
)

func calRemainderAndMod(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Invalid operation: Denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	numerator := 20
	denominator := 10 
	remainder, mod, err := calRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
}

 