// Shadowing variables:  Variable that has the same name as a variable in a containing block 

package main 

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 6
		fmt.Println(x)
	}
	fmt.Println(x)
}