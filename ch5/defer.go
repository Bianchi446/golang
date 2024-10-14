// In go, a cleanup code is attached to the function with the defer keyword 

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Ensure a file argument is passed
	if len(os.Args) < 2 {
		fmt.Println("No file specified")
		return
	}

	// Open the file specified as command-line argument
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Buffer to hold the file data
	data := make([]byte, 2048)

	// Read the file in chunks
	for {
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break // Exit the loop on EOF
		}
		// Write the read data to stdout
		os.Stdout.Write(data[:count])
	}
}
