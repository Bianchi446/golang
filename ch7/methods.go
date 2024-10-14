package main

import (
	"fmt"
)

// Define the 'person' struct
type person struct {
	firstName string
	lastName  string // Changed LastName to lastName for consistency
	age       int
}

// Define the String method for the 'person' type
func (p person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.firstName, p.lastName, p.age)
}

func main() {
	// Instantiate a 'person' and print it
	p := person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	// Call the String method and print the result
	fmt.Println(p.String())
}
