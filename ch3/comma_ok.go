/*
	Comma ok idiom to tell the difference between a key associated with zero value
	and a key that is not in the map 
*/

package main

import "fmt"

func main() {

	m := map[string]int{
		"Hello" : 0,
		"World" : 10,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["World"]
	fmt.Println(v, ok)

	v, ok = m["Goodbye"]
	fmt.Println(v, ok)
}