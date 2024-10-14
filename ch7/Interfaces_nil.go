// For an interface to be nil, both the type and the value must be nil
package main

import "fmt"

func main() {
	var s *string
	fmt.Println(s == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil)  // Not consider nil because it have a type: *string value: nil
}
