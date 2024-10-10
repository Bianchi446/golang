// Arrays are barely used in Go, slices are used instead. 

// Note 1 : You can't use a type conversion to convert arrays of different sizes to identical types

// Note 2 : You can' write a function with arrays of any size 

// Note 3 : You can't assign arrays of different sizes to the same variables 

package main

import "fmt"


// Slices are read using a bracket syntax

func slice(){

	x := make([]int, 5)

	x[3] = 20

	fmt.Println(x[3])
	
}

func main() {
	slice()
}