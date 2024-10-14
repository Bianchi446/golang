// Thinking in functions as data -- The implications of creating a closure that references 
// local variables and then passing the closure to another function. 

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define the person struct
	type person struct {
		FirstName string
		LastName  string
		Age       int
	}
	
	// Create a slice of persons
	persons := []person{
		{"Pat", "Anderson", 70},
		{"Alan", "Rodriguez", 30},
		{"Bruce", "Lopez", 10},
	}
	
	// Print the initial list of people
	fmt.Println("Initial:", persons)
	
	// Sort by last name
	sort.Slice(persons, func(i int, j int) bool {
		return persons[i].LastName < persons[j].LastName
	})
	fmt.Println("Sorted by last name:", persons)

	// Sort by age
	sort.Slice(persons, func(i int, j int) bool {
		return persons[i].Age < persons[j].Age
	})
	fmt.Println("Sorted by age:", persons)
}
