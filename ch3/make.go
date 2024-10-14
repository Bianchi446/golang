package main

import "fmt"


// Note: If you have an idea of how large your slices need to be, but don't now
// in advance what values are those, use make. 

func main(){

	x := make([]int, 20)

	// Slice with default values

	data := []int{2, 3, 8}

	fmt.Println(data, x)
}