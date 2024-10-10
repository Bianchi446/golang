package main

import "fmt"

func main() {

	// Type conversions

	var x int = 190
	var y float64 = 46.9
	var w float64 = float64(x) + y
	var z int = x + int(y)

	fmt.Println(x, y, w, z)
}
