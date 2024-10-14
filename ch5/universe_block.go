package main
import(
	"fmt"
	"errors"
)

func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}

	return numerator / denominator, numerator % denominator, nil
}

func main() {
	// Print the result with correct number of return values
	result, remainder, err := divAndRemainder(20, 10)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
	}
}