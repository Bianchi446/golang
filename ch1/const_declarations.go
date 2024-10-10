package main

import "fmt"

	const x int64 = 20

	const (
		idkey = "id"
		idName = "name"
	)

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}

