// Advice : either never use append with a sub-slice or make sure that append
// doesn't cause 

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 30)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
}