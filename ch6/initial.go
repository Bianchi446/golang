/*
	Initial considerations: 
	1. Pointer arithmetic is not allowed in time 
	2. The * operator is the indirection operator -- Preceeds a pointer type and returns a pointer value 
	3. Before dereferencing a pointer, the pointer must be non-nill
	4. A pointer type is a type that represents a pointer 
	5. If you have an struct with a field of a pointer to a primitive type, you can't assign a 
		literal directly to that field. 

	
		*/
package main

import(
	"fmt"
)

func main() {

	x := 10 
	pointerToX := &x 
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)

	// Pointer is not nil ?

	var x *int
	fmt.Println(x == nil)
	fmt.Println(*x)

	var x = new(int)
	var pointerToX *int
	pointerToX = &x

	// Structs 

	type person struct {
		firstName string
		MiddleName *string
		LastName string
 	}

	p := person {
		firstName : "Alan",
		MiddleName : "Leon", // Will not compile
		LastName : "Lopez",
	}

	
}


