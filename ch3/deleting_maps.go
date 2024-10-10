// Removing key-value pairs 

package main

import "fmt"

func main() {
	m := map[string]int {
		"Hello" : 10,
		"world" : 20,
	}

	delete(m, "hello")

	fmt.Println(m)
}