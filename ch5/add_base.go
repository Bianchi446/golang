package main

import "fmt"

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals)) // Initialize an output slice
	for _, v := range vals {
		out = append(out, base+v) // Remove shadowing by not redeclaring 'out'
	}

	return out
}

func main() {
	// Correcting the format specifier for Printf
	fmt.Println(addTo(1))                      // Should print an empty slice since no values are passed
	fmt.Println(addTo(3, 2))                   // Should print [5]
	fmt.Println(addTo(1, 2, 3, 4, 6))          // Should print [3 4 5 7]

	// Using correct case for 'A'
	A := []int{4, 3}
	fmt.Println(addTo(3, A...))                // Correctly pass a slice as variadic arguments
	
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // Correctly passing inline slice with variadic syntax
}
