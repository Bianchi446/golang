package main

import "fmt"

func main() {
	samples := []string{"Hello", "world", "How", "you ", "doing?"}
	for _, sample := range samples { 
	for i, r := range sample{
		fmt.Println(i, r, string(r))
	}
	fmt.Println()
	}
}