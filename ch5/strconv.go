package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Invalid expression", err)
	} else {
		fmt.Println("Converted string: ", num)
	}

	// Trying to convert an inverted integer 
	// Error "Invalid expressions strconv.Atoi: Parsing "abc": invalid syntax"

	num, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("Invalid expression", err)
	} else {
		fmt.Println("Converted string: ", num)
	}
}