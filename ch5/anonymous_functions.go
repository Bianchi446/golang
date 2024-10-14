package main 

import(
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j ," From inside anonymous function" )
		}(i)
	}
}