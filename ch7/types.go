/*
	Type assertions and type switches 

	Go provides a way to see if a variable of an interface type has a specific concrete type 
	or if the concrete type implements another interface. 

	Note: Type conversion is not type assertion; type conversion can be applied to concrete types
	and interfaces and run in runtime, type assertions can only be applied to interface types

	*/


// Type assertion: Names the concrete types that implemented the interface

package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)        // Remember, Go is careful about concrete types
	fmt.Println(i2 + 1)
}

	