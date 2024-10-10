package main

import "fmt"

func teams() {
	var nilMap map[string]int

	totalWins := map[string]int{}

	teams := map[string][]string {
		"Orcas" : []string {"Fred", "Paul"},
		"Dolhpins" : []string {"Jose", "Laura"},
		"Marmots" : []string {"Darius", "Alban"},
	}

	fmt.Println(teams, nilMap, totalWins)

	totalWins["Orcas"] = 1
	totalWins["Dolphins"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Marmots"])
	totalWins["Marmots"]++
	fmt.Println(totalWins["Marmots"])
	totalWins["Dolphins"] = 3
	fmt.Println(totalWins["Dolphins"])
}


func main() {
	teams()
}