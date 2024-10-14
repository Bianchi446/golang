/*
	The real star of Go's design is its implicit interfaces, the only abstract type in Go

	1. In an interface declaration, and interface literal appears after the name of the interface type

	2. Interfaces are usually named with 'er' endings: Handler, Reader, Stringer, Marshaler, Closer

	3. The application will change over time; you need flexibility (In terms of implementation)
		Also, you need to specify in what your code depends on in order for people to understand it

		Answer: Implicit interfaces
*/

// Explanation: Interface but only the client knows about it;



type logicProvider struct {}

func (lp logicProvider) Process(data string) string {
	// Business logic
}

type Logic interface {
	process(data string )
}

type Client struct {
	L logic
}

func(c client) Program() {
	c.L.Process(data)
}

main() {
	c := Client{
		L : logicProvider{},
	}
	c.Program()
}
