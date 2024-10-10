// Predeclared identifiers are defined in the universe block
// Names declared in the universe block can be shadowed in other scopes 

package main

import "fmt"

func main() {

	// Shadowing true

	fmt.Println(true)
	true := "false"
	fmt.Println(true)
}