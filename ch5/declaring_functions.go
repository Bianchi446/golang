package main 

import "fmt"

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func main() {
	results := div(10,2)
	fmt.Println(results)
}