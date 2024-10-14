/*
	Pointers make it harder to understand data flow and create extra work for the garbage collector 

	Note: Rather than populating a struct by passing a pointer to it into a function,
	have the function instantiate and return the struct.

*/

// Populating a struct by passing a pointer to it into a function

package main

import (
	"fmt"
)

// Define the Foo struct
type Foo struct {
	Field1 string
	Field2 int
}

// Function that modifies a passed-in pointer to Foo
func makeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

// Function that instantiates and returns a Foo struct
func makeFoo2() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20, // Corrected the typo from Filed2 to Field2
	}
	return f, nil
}

func main() {
	// Call makeFoo2() to create and return a new Foo struct
	foo, err := makeFoo2()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Foo from makeFoo2:", foo)

	// Create an empty Foo struct and pass its pointer to makeFoo()
	var fooPtr Foo
	err = makeFoo(&fooPtr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Foo after makeFoo:", fooPtr)
}
