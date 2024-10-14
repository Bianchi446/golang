package main 

import(
	"fmt"
)

type person struct {
	name string
	age int 
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	i := 2
	s := "Hello"
	p := person{}
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}