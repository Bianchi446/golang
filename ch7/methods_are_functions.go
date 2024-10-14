/*
	Methods in Go are like functions that can be used as a replacement
	for a function any times there's a variable or parameter of a function type


	Functions vs methods 

	If the logic depends on what values are configured at startup 
	or changed while the program is running, those values must be stored
	in a struct and the logic implemented as a method. 

	If the logic only depends on the input parameters, it should 
	be a function

	*/

package main

import(
	"fmt"
)

type Adder struct {
	start int 
}

func (a Adder) AddTo(val int) int {
	return a.start + val
} 

func main() {

	// Create an instance of the type and invoke its method 
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	f2 := myAdder.AddTo
	fmt.Println(f2(15))
}