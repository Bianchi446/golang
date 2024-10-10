package main

import "fmt"

func main() {
	
	type firstPerson struct {
		name string
		age int
		
	}

	f := firstPerson{
		name: "Bob",
		age: 70,
	}

	var g struct {
		name string
		age int 
	}

	f = g
	fmt.Println(f == g)
}