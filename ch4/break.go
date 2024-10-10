package main

import "fmt"

func main() {
	
	countdown := 10

	for {
		fmt.Println("Countdown: ", countdown)

		countdown--

		if countdown <= 0 {
			break
		}
	}

	fmt.Println("Lift off")
}

