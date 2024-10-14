// Final note: Don't write getter and setter methods for Go structs. 


package main 

import(
	"fmt"
	"time"
)


type Counter struct {
	total 	int 
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("Total: %d, last updated: %v", c.total, c.lastUpdated)
}	

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong: ", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight: ", c.String())
}



func main() {

	var c Counter
	doUpdateWrong(c)
	fmt.Println("In main: ", c.String())
	doUpdateRight(&c)
	fmt.Println("In main: ", c.String())
}
