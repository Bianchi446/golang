package main

import "fmt"

func main() {

	var x = []float64{1, 2, 3}

	x = append(x, 10.2, 11.2, 12.3)

	fmt.Println(x)
}