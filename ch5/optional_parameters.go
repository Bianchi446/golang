package main

import "fmt"

// Define the options structure
type MyFuncOpts struct {
	firstName string
	lastName  string
	age       int
	address   string
}

// Define the function that takes options and returns an error
func MyFunc(opts MyFuncOpts) error {
	fmt.Printf("First Name: %s, Last Name: %s, Age: %d, Address: %s\n", opts.firstName, opts.lastName, opts.age, opts.address)
	return nil
}

func main() {
	// Call the function with specific options
	MyFunc(MyFuncOpts{
		lastName: "joe",
		age:      89,
	})

	MyFunc(MyFuncOpts{
		firstName: "alan",
		age:       20,
	})
}