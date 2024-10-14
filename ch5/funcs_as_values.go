package main

import (
	"fmt"
	"strconv"
)

// Arithmetic functions
func add(j int, i int) int {
	return j + i
}

func substraction(j int, i int) int {
	return j - i
}

func multiplication(j int, i int) int {
	return j * i 
}

func division(j int, i int) int {
	if i == 0 {
		fmt.Println("Cannot divide by zero")
		return 0
	}
	return j / i
}

// Map of operations
var OperationsMap = map[string]func(int, int) int{
	"+": add,
	"-": substraction,
	"*": multiplication,
	"/": division,
}

func main() {
	// Sample expressions
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"two", "+", "three"},
		{"5"}, // Invalid case with missing operator/operand
	}

	// Process each expression
	for _, expression := range expressions {
		// Ensure the expression has exactly 3 parts
		if len(expression) != 3 {
			fmt.Println("Expression is not valid: ", expression)
			continue
		}

		// Convert the first operand
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println("Error converting the first operand:", err)
			continue
		}

		// Get the operator and corresponding function
		op := expression[1]
		opFunc, ok := OperationsMap[op]
		if !ok {
			fmt.Println("Unsupported operator:", op)
			continue
		}

		// Convert the second operand
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println("Error converting the second operand:", err)
			continue
		}

		// Perform the operation and print the result
		result := opFunc(p1, p2)
		fmt.Println("Result of", expression, "is", result)
	}
}
