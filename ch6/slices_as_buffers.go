/*
	Even though Go is a garbage-collector language, writting idiomatic Go means 
	avoiding unneded allocations. 

		We create a slice of bytes and use it as a buffer to read data from data source 
*/


package main

import (
	"fmt"
	"os"
	"io"
)

func process(data []byte) {
	// Placeholder function for processing data
	fmt.Println(string(data))
}

func main() {
	// Open the file "test.txt"
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Ensure the file is closed when we're done
	defer file.Close()

	// Create a buffer to read the data
	data := make([]byte, 100)

	// Read the file in a loop
	for {
		count, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				// Handle error, except for EOF which indicates end of file
				fmt.Println("Error reading file:", err)
				return
			}
			break
		}

		// Process the read data
		if count > 0 {
			process(data[:count])
		}
	}
}
