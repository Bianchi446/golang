// What if you have a nested for loops and you want to exit an iterator of an outer loop?

package main

import "fmt"

func main(){
	samples := []string{"Hello", "apple_n!"}
outer: 
	for _, sample := range samples {
		for i,r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}