package main 

import "fmt"

/*
	Maps featureS:
	- growth with key-value pairs
	- You can use make to create a map with specific initial size
	- The zero value of map is Nil
	- Maps are not comparable (== or !=)
*/

func main() {
	var nilMap map[string]int

	totalWins := map[string]int{}

	teams := map[string][]string {
		"Orcas" : []string {"Fred", "Paul"},
		"Dolhpins" : []string {"Jose", "Laura"},
		"Marmots" : []string {"Darius", "Alban"},
	}

	fmt.Println(teams, nilMap, totalWins)
}